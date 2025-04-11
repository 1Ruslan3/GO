package main

import (
	"fmt"
)

func main() {
	ages := map[string]int{
		"alice":  31,
		"mark":   29,
		"ruslan": 21,
	}

	ages["alice"] = 32
	ages["bob"] = 12

	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}

}
