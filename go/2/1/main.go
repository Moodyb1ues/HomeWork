package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Print("Введите число: ")
	fmt.Scan(&num) // читаем число от пользователя

	if num%2 == 0 {
		fmt.Println(num, "четное")
	} else {
		fmt.Println(num, "нечетное")
	}
}