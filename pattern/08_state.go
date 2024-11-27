package pattern

import "fmt"

/*
				Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
			https://en.wikipedia.org/wiki/State_pattern
*/

/*
	1) Состояние — это поведенческий паттерн проектирования, который позволяет объектам менять поведение в зависимости от
	своего состояния. Извне создаётся впечатление, что изменился класс объекта.
	2) Паттерн "Состояние" применим в следующих случаях:
	- Когда есть объект, поведение которого кардинально меняется в зависимости от внутреннего состояния,
	причём типов состояний много, и их код часто меняется.
	- Когда код класса содержит множество больших, похожих друг на друга, условных операторов,
	которые выбирают поведения в зависимости от текущих значений полей класса.
	- Когда вы сознательно используете табличную машину состояний, построенную на условных операторах,
	но вынуждены мириться с дублированием кода для похожих состояний и переходов.
	3) Плюсы:
	- Избавляет от множества больших условных операторов машины состояний;
	- Концентрирует в одном месте код, связанный с определённым состоянием;
	- Упрощает код контекста.
	4) Минусы:
	- Может неоправданно усложнить код, если состояний мало и они редко меняются.
	5) Реальные примеры применения:
	- В пакетах io и io/ioutil поведение объектов, реализующих интерфейсы Reader и Writer,
	может изменяться в зависимости от состояния потока данных или самого объекта;
	- Контексты в Go могут представлять разные состояния выполнения за счет своей вложенной структуры и механизма отмены.
	Разные реализации Context могут определять, как следует обрабатывать запросы в зависимости от того, был ли контекст отменен
	или достигнут таймаут.
	- Структуры синхронизации (например, sync.Mutex, sync.WaitGroup) меняют свое поведение в зависимости от внутреннего состояния
	(заблокировано или разблокировано).
*/

// Реализуем данный паттерн на примере торгового автомата с двумя состояниями

// Контекст
type VendingMachine struct {
	WaitChoice State
	PushItem   State

	currentState State
}

func newVendingMachine() *VendingMachine {
	m := &VendingMachine{}
	waitChoice := &StateWait{vendingMachine: m}
	pushItem := &StateGiving{vendingMachine: m}
	m.setState(waitChoice)
	m.PushItem = pushItem
	m.WaitChoice = waitChoice
	return m
}

func (v *VendingMachine) setState(s State) {
	v.currentState = s
}

// Интерфейс состояний
type State interface {
	requestItem()
	giveItem()
}

// Конкретное состояние 1
type StateWait struct {
	vendingMachine *VendingMachine
}

func (s *StateWait) requestItem() {
	fmt.Println("Item requested")
	s.vendingMachine.setState(s.vendingMachine.PushItem)
}

func (s *StateWait) giveItem() {
	fmt.Println("Before you must request item!!!")
}

// Конкретное состояние 2
type StateGiving struct {
	vendingMachine *VendingMachine
}

func (s *StateGiving) requestItem() {
	fmt.Println("Before you must take item!!!")
}

func (s *StateGiving) giveItem() {
	fmt.Println("Item given")
	s.vendingMachine.setState(s.vendingMachine.WaitChoice)
}

// Клиентский код
func CheckState() {
	machine := newVendingMachine()

	machine.currentState.giveItem()
	machine.currentState.requestItem()

	machine.currentState.requestItem()
	machine.currentState.giveItem()
	machine.currentState.requestItem()
}

/*
	Вывод:
Before you must request item!!!
Item requested
Before you must take item!!!
Item given
Item requested
*/
