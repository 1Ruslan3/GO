package main

import (
	"fmt"
)

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

func main() {
	var a **int
	b := 6
	c := &b
	a = &c
	fmt.Println(a)

}
