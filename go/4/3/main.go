package main

import "fmt"

func main() {
	people := map[string]int{
		"Аня":  20,
		"Игорь": 25,
		"Саша": 22,
	}

	var name string
	fmt.Print("Введите имя для удаления: ")
	fmt.Scan(&name)

	delete(people, name)

	fmt.Println("Обновлённый список:")
	for k, v := range people {
		fmt.Printf("%s: %d\n", k, v)
	}
}