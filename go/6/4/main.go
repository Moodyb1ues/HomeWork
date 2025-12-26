package main

import (
	"fmt"
	"strings"
)

func main() {
	var text string
	fmt.Print("Введите строку: ")
	fmt.Scanln(&text)

	fmt.Println("Верхний регистр:", strings.ToUpper(text))
}