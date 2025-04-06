package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Node represents a node in the arithmetic expression tree
type Node struct {
	Value string
	Left  *Node
	Right *Node
}

// NewNode creates a new node with the given value
func NewNode(value string) *Node {
	return &Node{Value: value}
}

// IsOperator checks if the node's value is an operator
func (n *Node) IsOperator() bool {
	return n.Value == "+" || n.Value == "-" || n.Value == "*" || n.Value == "/"
}

// Evaluate computes the result of the arithmetic expression
func (n *Node) Evaluate() (float64, error) {
	if !n.IsOperator() {
		// If it's a number, convert it to float64
		return strconv.ParseFloat(n.Value, 64)
	}

	// Evaluate left and right subtrees
	left, err := n.Left.Evaluate()
	if err != nil {
		return 0, err
	}

	right, err := n.Right.Evaluate()
	if err != nil {
		return 0, err
	}

	// Perform the operation
	switch n.Value {
	case "+":
		return left + right, nil
	case "-":
		return left - right, nil
	case "*":
		return left * right, nil
	case "/":
		if right == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return left / right, nil
	default:
		return 0, fmt.Errorf("unknown operator: %s", n.Value)
	}
}

// String returns the string representation of the expression
func (n *Node) String() string {
	if !n.IsOperator() {
		return n.Value
	}
	return fmt.Sprintf("(%s %s %s)", n.Left.String(), n.Value, n.Right.String())
}

// ParseExpression parses a string expression and builds an expression tree
func ParseExpression(expr string) (*Node, error) {
	// Remove all whitespace
	expr = strings.ReplaceAll(expr, " ", "")

	// Find the operator with lowest precedence
	lowestPrec := -1
	lowestPos := -1
	parenCount := 0

	for i := len(expr) - 1; i >= 0; i-- {
		char := string(expr[i])
		if char == ")" {
			parenCount++
		} else if char == "(" {
			parenCount--
		} else if parenCount == 0 && (char == "+" || char == "-") {
			lowestPos = i
			lowestPrec = 1
			break
		} else if parenCount == 0 && (char == "*" || char == "/") && lowestPrec < 2 {
			lowestPos = i
			lowestPrec = 2
		}
	}

	if lowestPos == -1 {
		// Если оператор не найден, проверяем является ли выражение числом или выражением в скобках
		if expr[0] == '(' && expr[len(expr)-1] == ')' {
			// Если выражение в скобках - рекурсивно разбираем его содержимое без скобок
			return ParseExpression(expr[1 : len(expr)-1])
		}
		// Если это не выражение в скобках, значит это число - создаем для него узел
		return NewNode(expr), nil
	}

	// Create a node for the operator
	node := NewNode(string(expr[lowestPos]))

	// Parse left and right subtrees
	left, err := ParseExpression(expr[:lowestPos])
	if err != nil {
		return nil, err
	}

	right, err := ParseExpression(expr[lowestPos+1:])
	if err != nil {
		return nil, err
	}

	node.Left = left
	node.Right = right

	return node, nil
}

func main() {
	// Example expressions
	expressions := []string{
		"3 + 4 * 5",
		"(3 + 4) * (5 - 2)",
		"10 / 2 + 3 * 4",
		"1 + 2 + 3 + 4",
	}

	for _, expr := range expressions {
		fmt.Printf("\nExpression: %s\n", expr)

		// Parse and build the tree
		root, err := ParseExpression(expr)
		if err != nil {
			fmt.Printf("Error parsing expression: %v\n", err)
			continue
		}

		// Print the expression with parentheses
		fmt.Printf("Parsed expression: %s\n", root.String())

		// Evaluate the expression
		result, err := root.Evaluate()
		if err != nil {
			fmt.Printf("Error evaluating expression: %v\n", err)
			continue
		}
		fmt.Printf("Result: %.2f\n", result)
	}
}
