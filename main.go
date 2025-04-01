package main

import "fmt"

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
	rect := Rectangle{Width: 2, Height: 3}
	fmt.Println(rect.Area())
	rect.Scale(2)
	fmt.Println(rect.Area())
}
