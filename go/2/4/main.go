package main

import (
	"fmt"
)

func stringLength(s string) int {
	return len(s)
}

func main() {
	var text string
	fmt.Print("Введите строку: ")
	fmt.Scanln(&text) // читает строку до пробела
	// если нужно читать с пробелами -> заменить на fmt.Scanf("%[^\n]", &text)

	fmt.Println("Длина строки:", stringLength(text))
}