package main

import (
	"fmt"
	"time"
)

func calculateFloatOperations(a, b float64) (float64, float64) {
	sum := a + b
	difference := a - b
	return sum, difference
}

func main() {
	fmt.Println("1. Текущее время и дата:")
	currentTime := time.Now()
	fmt.Printf("   Время: %s\n", currentTime.Format("15:04:05"))
	fmt.Printf("   Дата: %s\n", currentTime.Format("02.01.2006"))
	fmt.Printf("   Полная дата и время: %s\n\n", currentTime.Format("02.01.2006 15:04:05"))

	fmt.Println("2. Переменные различных типов:") 
	var intVar int = 42
	var floatVar float64 = 3.14159
	var stringVar string = "поставьте пять"
	var boolVar bool = true
	fmt.Printf("   int: %d\n", intVar)
	fmt.Printf("   float64: %.5f\n", floatVar)
	fmt.Printf("   string: %s\n", stringVar)
	fmt.Printf("   bool: %t\n\n", boolVar)

	fmt.Println("3. Краткая форма объявления переменных:")
	name := "Alice"
	age := 25
	height := 1.75
	isStudent := true
	fmt.Printf("   Имя: %s (тип: %T)\n", name, name)
	fmt.Printf("   Возраст: %d (тип: %T)\n", age, age)
	fmt.Printf("   Рост: %.2f (тип: %T)\n", height, height)
	fmt.Printf("   Студент: %t (тип: %T)\n\n", isStudent, isStudent)

	fmt.Println("4. Арифметические операции с целыми числами:")
	num1 := 20
	num2 := 8
	fmt.Printf("   Числа: %d и %d\n", num1, num2)
	fmt.Printf("   Сумма: %d + %d = %d\n", num1, num2, num1+num2)
	fmt.Printf("   Разность: %d - %d = %d\n", num1, num2, num1-num2)
	fmt.Printf("   Произведение: %d * %d = %d\n", num1, num2, num1*num2)
	fmt.Printf("   Деление: %d / %d = %d\n", num1, num2, num1/num2)
	fmt.Printf("   Остаток от деления: %d %% %d = %d\n\n", num1, num2, num1%num2)

	fmt.Println("5. Сумма и разность чисел с плавающей запятой:")
	float1 := 15.7
	float2 := 4.3
	sum, diff := calculateFloatOperations(float1, float2)
	fmt.Printf("   Числа: %.2f и %.2f\n", float1, float2)
	fmt.Printf("   Сумма: %.2f + %.2f = %.2f\n", float1, float2, sum)
	fmt.Printf("   Разность: %.2f - %.2f = %.2f\n\n", float1, float2, diff)

	fmt.Println("6. Среднее значение трех чисел:")
	a := 12.5
	b := 18.3
	c := 22.1
	average := (a + b + c) / 3
	fmt.Printf("   Числа: %.2f, %.2f, %.2f\n", a, b, c)
	fmt.Printf("   Среднее: (%.2f + %.2f + %.2f) / 3 = %.2f\n", a, b, c, average)
}
