package pattern

import (
	"fmt"
)

/*
				Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
		https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

/*
	1) Цепочка обязанностей — это поведенческий паттерн проектирования, который позволяет передавать запросы последовательно по цепочке
	обработчиков. Каждый последующий обработчик решает, может ли он обработать запрос сам и стоит ли передавать запрос дальше по цепи.
	2) Паттерн "Цепочка вызовов" применим в следующих случаях:
	- Когда программа должна обрабатывать разнообразные запросы несколькими способами, но заранее неизвестно,
	какие конкретно запросы будут приходить и какие обработчики для них понадобятся;
	- Когда важно, чтобы обработчики выполнялись один за другим в строгом порядке;
	- Когда набор объектов, способных обработать запрос, должен задаваться динамически.
	3) Плюсы:
	- Уменьшает зависимость между клиентом и обработчиками;
	- Реализует принцип единственной обязанности;
	- Реализует принцип открытости/закрытости.
	4) Минусы:
	- Запрос может остаться никем не обработанным.
	5) Реальный пример использования может быть связан с обработкой запроса доступа к какой-либо системе и обработкой HTTP запрососов.
*/

/*
	Реализуем данный паттерн на примере посещения больницы пациентом.
*/

// Пациент
type Patient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
}

// Интерфейс обработчика
type Departament interface {
	execute(*Patient)
	setNext(Departament)
}

// Конкретный обработчик
type Reception struct {
	next Departament
}

func (r *Reception) execute(p *Patient) {
	if p.registrationDone {
		fmt.Printf("Patient [%s] registration already done\n", p.name)
		r.next.execute(p)
		return
	}
	fmt.Printf("Reception registering patient [%s]\n", p.name)
	p.registrationDone = true
	r.next.execute(p)
}

func (r *Reception) setNext(d Departament) {
	r.next = d
}

// Конкретный обработчик 2
type Doctor struct {
	next Departament
}

func (d *Doctor) execute(p *Patient) {
	if p.doctorCheckUpDone {
		fmt.Printf("Doctor checkup already done for patient [%s]\n", p.name)
		d.next.execute(p)
		return
	}
	fmt.Printf("Doctor checking patient [%s]\n", p.name)
	p.doctorCheckUpDone = true
	d.next.execute(p)
}

func (doc *Doctor) setNext(d Departament) {
	doc.next = d
}

// Конкретный обработчик 3
type Medical struct {
	next Departament
}

func (m *Medical) execute(p *Patient) {
	if p.medicineDone {
		fmt.Printf("Medicine already given to patient [%s]\n", p.name)
		return
	}
	fmt.Printf("Medical giving medicine to patient [%s]\n", p.name)
	p.doctorCheckUpDone = true
}

func (m *Medical) setNext(d Departament) {
	m.next = d
}

// Клиентский код
func CkeckChainOfResp() {
	patient1 := &Patient{
		name:              "Robert",
		registrationDone:  false,
		doctorCheckUpDone: false,
		medicineDone:      false,
	}

	patient2 := &Patient{
		name:              "Jhon",
		registrationDone:  true,
		doctorCheckUpDone: true,
		medicineDone:      false,
	}

	medical := &Medical{}

	doctor := &Doctor{}
	doctor.setNext(medical)

	reception := &Reception{}
	reception.setNext(doctor)

	reception.execute(patient1)

	fmt.Println()

	doctor.execute(patient2)
}
