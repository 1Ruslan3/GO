package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

func main() {
	var dilbert Employee
	dilbert.Salary -= 5000
	fmt.Println(dilbert.Salary)

	dilbert.Position = "Plus"
	position := &dilbert.Position
	*position = "Senior " + *position
	fmt.Println(*position)

}
