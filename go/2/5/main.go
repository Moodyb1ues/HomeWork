package main

import (
	"fmt"
)

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	var w, h float64
	fmt.Print("Введите ширину и высоту прямоугольника: ")
	fmt.Scan(&w, &h)

	rect := Rectangle{Width: w, Height: h}
	fmt.Println("Площадь прямоугольника:", rect.Area())
}