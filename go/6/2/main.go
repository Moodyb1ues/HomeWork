package main

import "fmt"

func averageAge(m map[string]int) float64 {
	sum := 0
	for _, age := range m {
		sum += age
	}
	return float64(sum) / float64(len(m))
}

func main() {
	people := map[string]int{
		"Аня":  20,
		"Игорь": 25,
		"Саша": 22,
	}
	fmt.Println("Средний возраст:", averageAge(people))
}