package pattern

import (
	"errors"
	"fmt"
	"time"
)

/*
				Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
			https://en.wikipedia.org/wiki/Facade_pattern
*/

/*
	1) Фасад — это структурный паттерн проектирования, который предоставляет простой интерфейс к сложной системе классов, библиотеке или фреймворку.
	2) Паттерн фасад применим в следующих случаях:
	- Когда нужно представить простой или урезанный интерфейс к сложной подсистеме;
	- Когда нужно разложить подсистему на отдельные слои.
	3) Плюсы:
	- Использование паттерна позволяет не думать о внутреннем устройсве подсистемы.
	- Изолирует клиентa от компонентов сложной подсистемы.
	4) Минусы:
	- Фасад рискует стать божественным объектом, привязанным ко всем классам программы.
	- Как правило, имеет ограниченную функциональность в отличие от непосредственной работы с компонентами системы.
	5) Реальные примеры могут быть связаны с использованием различных API и библиотек. Также хорошим примером является фасад к базе данных,
	который упрощает работу с базой данных и скрывает сложность внутренней подсистемы.
*/

/*
	Реализуем паттерн Фасад на примере покупки товара в магазине.
*/

// Товар
type Product struct {
	Name  string
	Price float64
}

// Магазин (Объект фасада)
type Shop struct {
	Name     string
	Products []Product
}

// Реализация простого интерфейса (фасад)
func (shop Shop) Sell(user User, product string) error {
	fmt.Println("[Магазин] Запрос к пользователю, для получения остатка по карте")
	time.Sleep(time.Millisecond * 500)
	err := user.Card.CheckBalance()
	if err != nil {
		return err
	}

	fmt.Printf("[Магазин] Проверка может ли [%s] купить товар\n", user.Name)
	time.Sleep(time.Millisecond * 500)

	for _, prod := range shop.Products {
		if product != prod.Name {
			continue
		}
		if prod.Price > user.GetBalance() {
			return errors.New("У пользователя недостаточно средств для покупки товара")
		}
		fmt.Printf("[Магазин] Товар [%s] куплен\n", prod.Name)
	}
	return nil
}

// Карта
type Card struct {
	Name    string
	Balance float64
	Bank    *Bank
}

// Проверка баланса карты
func (card Card) CheckBalance() error {
	fmt.Println("[Карта] Запрос в банк для проверки остатка")
	time.Sleep(time.Millisecond * 800)
	return card.Bank.CheckBalance(card.Name)
}

// Банк
type Bank struct {
	Name  string
	Cards []Card
}

// Проверка баланса карты через банк
func (bank Bank) CheckBalance(cardNumber string) error {
	fmt.Printf("[Банк] Получение остатка по карте %s\n", cardNumber)
	time.Sleep(time.Millisecond * 300)

	for _, card := range bank.Cards {
		if card.Name != cardNumber {
			continue
		}
		if card.Balance <= 0 {
			return errors.New("[Банк] Недостаточно средств!")
		}
	}

	fmt.Println("[Банк] Остаток положительный!")
	return nil
}

// Покупатель
type User struct {
	Name string
	Card *Card
}

// Получение баланса
func (user User) GetBalance() float64 {
	return user.Card.Balance
}

// Инициализаия
var (
	bank = Bank{
		Name:  "BANK",
		Cards: []Card{},
	}
	card1 = Card{
		Name:    "CARD1",
		Balance: 500,
		Bank:    &bank,
	}
	card2 = Card{
		Name:    "CARD2",
		Balance: 5,
		Bank:    &bank,
	}
	user1 = User{
		Name: "Покупатель-1",
		Card: &card1,
	}
	user2 = User{
		Name: "Покупатель-2",
		Card: &card2,
	}
	prod = Product{
		Name:  "Колбаса",
		Price: 300,
	}
	shop = Shop{
		Name: "SHOP",
		Products: []Product{
			prod,
		},
	}
)

func CheckFacade() {
	bank.Cards = append(bank.Cards, card1, card2)
	// Ипользуем паттерн - получаем простой интерфейс
	err := shop.Sell(user1, prod.Name)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("[Продажа - 1] - OK")
	}

	// Ипользуем паттерн - получаем простой интерфейс
	err = shop.Sell(user2, prod.Name)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("[Продажа - 2] - OK")
	}
}

/*
	Вывод:
[Магазин] Запрос к пользователю, для получения остатка по карте
[Карта] Запрос в банк для проверки остатка
[Банк] Получение остатка по карте CARD1
[Банк] Остаток положительный!
[Магазин] Проверка может ли [Покупатель-1] купить товар
[Магазин] Товар [Колбаса] куплен
[Продажа - 1] - OK
[Магазин] Запрос к пользователю, для получения остатка по карте
[Карта] Запрос в банк для проверки остатка
[Банк] Получение остатка по карте CARD2
[Банк] Остаток положительный!
[Магазин] Проверка может ли [Покупатель-2] купить товар
У пользователя недостаточно средств для покупки товара
*/
