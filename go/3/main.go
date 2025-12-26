package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"practics/3/mathutils"
	"practics/3/stringutils"
)

func arrayExample() {
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println("массив:", arr)
}

func sliceExample() {
	arr := [5]int{10, 20, 30, 40, 50}
	s := arr[1:4]
	fmt.Println("исходный срез:", s)

	s = append(s, 60, 70)
	fmt.Println("после добавления:", s)

	if len(s) > 1 {
		s = append(s[:1], s[2:]...)
	}
	fmt.Println("после удаления:", s)
}

func longestStringExample() {
	words := []string{"Go", "Golang", "Программирование", "код"}
	longest := words[0]
	for _, w := range words {
		if len([]rune(w)) > len([]rune(longest)) {
			longest = w
		}
	}
	fmt.Println("самая длинная строка:", longest)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("введите число для факториала: ")
	var n int
	fmt.Fscan(reader, &n)
	_, _ = reader.ReadString('\n')
	fmt.Println("факториал =", mathutils.Factorial(n))

	fmt.Print("введите строку для переворота: ")
	line, _ := reader.ReadString('\n')
	line = strings.TrimRight(line, "\r\n")
	fmt.Println("перевернутая строка:", stringutils.Reverse(line))

	arrayExample()
	sliceExample()
	longestStringExample()
}
