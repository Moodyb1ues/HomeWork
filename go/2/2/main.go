package main

import (
	"fmt"
)

func checkNumber(n int) string {
	if n > 0 {
		return "Positive"
	} else if n < 0 {
		return "Negative"
	}
	return "Zero"
}

func main() {
	var num int
	fmt.Print("Введите число: ")
	fmt.Scan(&num)
	fmt.Println(checkNumber(num))
}