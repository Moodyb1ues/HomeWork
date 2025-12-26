package main

import (
	"fmt"
)

func average(a int, b int) float64 {
	return float64(a+b) / 2
}

func main() {
	var a, b int
	fmt.Print("Введите два числа: ")
	fmt.Scan(&a, &b)

	fmt.Println("Среднее значение:", average(a, b))
}