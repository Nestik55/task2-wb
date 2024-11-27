package pattern

import "fmt"

/*
				Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
			https://en.wikipedia.org/wiki/Visitor_pattern
*/

/*
	1) Посетитель — это поведенческий паттерн проектирования, который позволяет добавлять в программу
	новые операции, не изменяя классы объектов, над которыми эти операции могут выполняться.
	2) Паттерн Посетитель применим в следующих случаях:
	- Когда нужно выполнить какую-то операцию над всеми элементами сложной структуры объектов,
	например, деревом;
	- Когда над объектами сложной структуры объектов нужно выполнять некоторые не связанные между собой
	операции, но «засорять» классы такими операциями не хочется;
	- Когда новое поведение имеет смысл только для некоторых классов из существующей иерархии.
	3) Плюсы:
	- Упрощает добавление операций, работающих со сложными структурами объектов;
	- Объединяет родственные операции в одном классе;
	- Посетитель может накапливать состояние при обходе структуры элементов.
	4) Минусы:
	- Паттерн не оправдан, если иерархия элементов часто меняется;
	- Может привести к нарушению инкапсуляции элементов.
	5) Реальный примеры могут быть связаны с итерацией по структурам данных, подобным контейнеру. Также, например,
	визуализационные библиотеки могут использовать посетителя для обхода графических объектов и выполнения
	операций рендеринга или обработки событий.
*/

/*
	Реализуем данный паттерн на примере с различными геометрическими фигурами.
*/

// Элемент
type Shape interface {
	getType() string
	accept(Visitor)
}

// Конткретный элемент 1
type Square struct {
	side int
}

func (s *Square) getType() string {
	return "Square"
}

func (s *Square) accept(v Visitor) {
	v.visitForSquare(s)
}

// Конткретный элемент 2
type Circle struct {
	radius int
}

func (c *Circle) getType() string {
	return "Circle"
}

func (c *Circle) accept(v Visitor) {
	v.visitForCircle(c)
}

// Конткретный элемент 3
type Rectangle struct {
	a int
	b int
}

func (r *Rectangle) getType() string {
	return "Rectangle"
}

func (r *Rectangle) accept(v Visitor) {
	v.visitForRectangle(r)
}

// Посетитель
type Visitor interface {
	visitForSquare(*Square)
	visitForCircle(*Circle)
	visitForRectangle(*Rectangle)
}

// Конткретный посетитель 1
type AreaCalculator struct {
	//area int
}

func (a *AreaCalculator) visitForSquare(s *Square) {
	fmt.Println("Вычисление площади квадрата...")
}

func (a *AreaCalculator) visitForRectangle(r *Rectangle) {
	fmt.Println("Вычисление площади прямоугольника...")
}

func (a *AreaCalculator) visitForCircle(с *Circle) {
	fmt.Println("Вычисление площади круга...")
}

// Конткретный посетитель 2
type MidCordsCalculator struct {
	//area int
}

func (m *MidCordsCalculator) visitForSquare(s *Square) {
	fmt.Println("Вычисление центра квадрата...")
}

func (m *MidCordsCalculator) visitForRectangle(r *Rectangle) {
	fmt.Println("Вычисление центра прямоугольника...")
}

func (m *MidCordsCalculator) visitForCircle(с *Circle) {
	fmt.Println("Вычисление центра круга...")
}

// Клиентский код
func CheckVisitor() {
	square := &Square{side: 4}
	circle := &Circle{radius: 5}
	rectangle := &Rectangle{
		a: 4,
		b: 5,
	}

	areaCalculator := &AreaCalculator{}

	square.accept(areaCalculator)
	circle.accept(areaCalculator)
	rectangle.accept(areaCalculator)

	fmt.Println()

	midCordsCalculator := &MidCordsCalculator{}

	square.accept(midCordsCalculator)
	circle.accept(midCordsCalculator)
	rectangle.accept(midCordsCalculator)
}
