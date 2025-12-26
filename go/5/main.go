package main

import (
	"fmt"
	"math"
)

// ===== 1. структура Person =====
type Person struct {
	name string
	age  int
}

// метод для вывода информации о человеке
func (p Person) Info() {
	fmt.Printf("Имя: %s, Возраст: %d\n", p.name, p.age)
}

// ===== 2. метод birthday =====
func (p *Person) Birthday() {
	p.age++
}

// ===== 3. структура Circle =====
type Circle struct {
	radius float64
}

// метод для вычисления площади круга
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// ===== 4. интерфейс Shape =====
type Shape interface {
	Area() float64
}

// структура Rectangle
type Rectangle struct {
	width  float64
	height float64
}

// реализация Area() для Rectangle
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// ===== 5. функция для работы со срезом Shape =====
func PrintAreas(shapes []Shape) {
	for _, s := range shapes {
		fmt.Println("Площадь =", s.Area())
	}
}

// ===== 6. интерфейс Stringer =====
type Stringer interface {
	String() string
}

// структура Book
type Book struct {
	title  string
	author string
	pages  int
}

// реализация String() для Book
func (b Book) String() string {
	return fmt.Sprintf("Книга: \"%s\", Автор: %s, Страниц: %d", b.title, b.author, b.pages)
}

func main() {
	// 1–2. работа с Person
	p := Person{name: "Саша", age: 20}
	p.Info()
	p.Birthday()
	p.Info()

	// 3–5. работа с Circle, Rectangle и интерфейсом Shape
	c := Circle{radius: 5}
	r := Rectangle{width: 4, height: 6}
	shapes := []Shape{c, r}
	PrintAreas(shapes)

	// 6. работа с Book и интерфейсом Stringer
	b := Book{title: "Приключения", author: "Иванов", pages: 300}
	var s Stringer = b
	fmt.Println(s.String())
}
