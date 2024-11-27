package pattern

import "fmt"

/*
				Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
			https://en.wikipedia.org/wiki/Builder_pattern
*/

/*
	1) Строитель — это порождающий паттерн проектирования, который позволяет создавать сложные объекты пошагово.
	Строитель даёт возможность использовать один и тот же код строительства для получения разных представлений объектов.
	2) Паттерн Строитель применим в следующих случаях:
	- Когда нужно избавиться от «телескопического конструктора»;
	- Когда код должен создавать разные представления какого-то объекта.
	- Когда нужно собирать сложные составные объекты.
	3) Плюсы:
	- Позволяет создавать продукты пошагово.
	- Позволяет использовать один и тот же код для создания различных продуктов.
	- Изолирует сложный код сборки продукта от его основной бизнес-логики.
	4) Минусы:
	- Усложняет код программы из-за введения дополнительных классов.
	- Клиент будет привязан к конкретным классам строителей, так как в интерфейсе директора может не быть метода получения результата.
	5) Реальные примеры использования данного паттерна можно найти в различных библиотеках и фреймворках.
	Например, библиотека "sqlx" использует билдер для построения SQL-запросов с помощью цепочки методов.
	Билдер позволяет создавать сложные SQL-запросы, добавляя условия, сортировку, ограничения и другие параметры.
*/

/*
	Реализуем данный паттерн на примере создания объктов домов.
*/

// Продукт (дом)
type House struct {
	windowType string
	doorType   string
	floor      int
}

// Интерфейс cтроителя
type IBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() House
}

func getBuilder(builderType string) IBuilder {
	switch builderType {
	case "normal":
		return newNormalBulder()
	case "igloo":
		return newIglooBulder()
	default:
		return nil
	}
}

// Конкретный строитель 1
type NormalBulder struct {
	windowType string
	doorType   string
	floor      int
}

func newNormalBulder() *NormalBulder {
	return &NormalBulder{}
}

func (b *NormalBulder) setWindowType() {
	b.windowType = "Wooden window"
}

func (b *NormalBulder) setDoorType() {
	b.doorType = "Wooden door"
}

func (b *NormalBulder) setNumFloor() {
	b.floor = 2
}

func (b *NormalBulder) getHouse() House {
	return House{
		windowType: b.doorType,
		doorType:   b.doorType,
		floor:      b.floor,
	}
}

// Конкретный строитель 2
type IglooBulder struct {
	windowType string
	doorType   string
	floor      int
}

func newIglooBulder() *IglooBulder {
	return &IglooBulder{}
}

func (b *IglooBulder) setWindowType() {
	b.windowType = "Snow window"
}

func (b *IglooBulder) setDoorType() {
	b.doorType = "Snow door"
}

func (b *IglooBulder) setNumFloor() {
	b.floor = 1
}

func (b *IglooBulder) getHouse() House {
	return House{
		windowType: b.doorType,
		doorType:   b.doorType,
		floor:      b.floor,
	}
}

// Директор
type Director struct {
	builder IBuilder
}

func newDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) setBuilder(b IBuilder) {
	d.builder = b
}

func (d *Director) buildHouse() House {
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}

// Клиентский код
func CheckBuilder() {
	normalBuilder := getBuilder("normal")
	iglooBuilder := getBuilder("igloo")

	director := newDirector(normalBuilder)
	normalHouse := director.buildHouse()

	fmt.Printf("Normal House Door Type: %s\n", normalHouse.doorType)
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.windowType)
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.floor)

	director.setBuilder(iglooBuilder)
	iglooHouse := director.buildHouse()

	fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.doorType)
	fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.windowType)
	fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.floor)

}
