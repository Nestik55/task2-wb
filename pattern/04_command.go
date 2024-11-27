package pattern

import "fmt"

/*
				Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
			https://en.wikipedia.org/wiki/Command_pattern
*/

/*
	1) Команда — это поведенческий паттерн проектирования, который превращает запросы в объекты, позволяя передавать их как аргументы
	при вызове методов, ставить запросы в очередь, логировать их, а также поддерживать отмену операций.
	2) Паттерн команда применим в следующтх случаях:
	- Когда нужно параметризовать объекты выполняемым действием;
	- Когда нужно ставить операции в очередь, выполнять их по расписанию или передавать по сети;
	- Когда нужна операция отмены.
	3) Плюсы:
	- Убирает прямую зависимость между объектами, вызывающими операции, и объектами, которые их непосредственно выполняют;
	- Позволяет реализовать простую отмену и повтор операций;
	- Позволяет реализовать отложенный запуск операций;
	- Позволяет собирать сложные команды из простых;
	- Реализует принцип открытости/закрытости.
	4) Минусы:
	- Усложняет код программы из-за введения множества дополнительных классов.
	5) Реальный примеры использования могут быть связаны с кнопками пользовательского интерфейса и пунктами меню,
	записью макросов, многоуровневой отменой операций (Undo), транзакциями и пулами потоков.
*/

/*
	Реализуем данный паттерн на примере кнопки включения/выключения устройства.
*/

// Отправитель
type Button struct {
	command Command
}

func (b *Button) press() {
	b.command.execute()
}

// Интерфейс команды
type Command interface {
	execute()
}

// Конткретная команда 1
type OnCommand struct {
	device Device
}

func (c *OnCommand) execute() {
	c.device.on()
}

// Конкретная команда 2
type OffCommand struct {
	device Device
}

func (c *OffCommand) execute() {
	c.device.off()
}

// Интерфейс получателя (устройства)
type Device interface {
	on()
	off()
}

// Конкретное конкретный получатель (устройство)
type Phone struct {
	isRunning bool
}

func (p *Phone) on() {
	p.isRunning = true
	fmt.Println("Phone is working")
}

func (p *Phone) off() {
	p.isRunning = false
	fmt.Println("Phone is not working")
}

// Клиентский код
func CheckCommand() {
	phone := &Phone{}
	onCommand := &OnCommand{device: phone}
	offCommand := &OffCommand{device: phone}

	onButton := &Button{
		command: onCommand,
	}
	onButton.press()

	onButton2 := &Button{
		command: onCommand,
	}

	offButton := &Button{
		command: offCommand,
	}
	offButton.press()

	onButton2.press()

}

/*
		Вывод:
	Phone is working
	Phone is not working
	Phone is working
*/
