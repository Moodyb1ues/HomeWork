package main

import "fmt"

func main() {
	people := map[string]int{
		"Аня":  20,
		"Игорь": 25,
		"Саша": 22,
	}

	// добавляем нового человека
	people["Мария"] = 30

	// выводим всех
	fmt.Println("Список людей и возрастов:")
	for name, age := range people {
		fmt.Printf("%s: %d\n", name, age)
	}
}