package main

import "fmt"

func main() {
	var n int
	fmt.Print("Сколько чисел хотите ввести? ")
	fmt.Scan(&n)

	sum := 0
	for i := 0; i < n; i++ {
		var x int
		fmt.Print("Введите число: ")
		fmt.Scan(&x)
		sum += x
	}

	fmt.Println("Сумма =", sum)
}