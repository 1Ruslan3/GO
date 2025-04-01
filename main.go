package main

import "fmt"

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 { // Метод по значению
	return r.Width * r.Height
}

func (r *Rectangle) Scale(factor float64) { // Метод по указателю
	r.Width *= factor
	r.Height *= factor
}

func main() {
	rect := Rectangle{Width: 2, Height: 3}
	fmt.Println(rect.Area()) // Вывод: 6
	rect.Scale(2)
	fmt.Println(rect.Area()) // Вывод: 12
}
