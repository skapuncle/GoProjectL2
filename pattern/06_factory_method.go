package main

/* Паттерн "Фабричный метод" (Factory Method) - это порождающий паттерн проектирования,
который определяет общий интерфейс для создания объектов в суперклассе,
позволяя подклассам изменять тип создаваемых объектов.

Преимущества паттерна "Фабричный метод" включают:

Избегание прямых связей между клиентом и конкретными классами продуктов.
Возможность расширять систему без изменения кода клиентов.

Минусы паттерна "Фабричный метод" включают:

Увеличение сложности кода из-за добавления новых классов и интерфейсов.
Неудобство при работе с большим количеством различных типов продуктов.
*/

import "fmt"

// Product представляет интерфейс для продукта.
type Product interface {
	Use()
}

// ConcreteProductA реализует интерфейс Product.
type ConcreteProductA struct{}

func (p *ConcreteProductA) Use() {
	fmt.Println("Using product A")
}

// ConcreteProductB реализует интерфейс Product.
type ConcreteProductB struct{}

func (p *ConcreteProductB) Use() {
	fmt.Println("Using product B")
}

// Creator представляет интерфейс для создателя.
type Creator interface {
	CreateProduct() Product
}

// ConcreteCreatorA реализует интерфейс Creator.
type ConcreteCreatorA struct{}

func (cc *ConcreteCreatorA) CreateProduct() Product {
	return new(ConcreteProductA)
}

// ConcreteCreatorB реализует интерфейс Creator.
type ConcreteCreatorB struct{}

func (cc *ConcreteCreatorB) CreateProduct() Product {
	return new(ConcreteProductB)
}

func main() {
	creators := []Creator{&ConcreteCreatorA{}, &ConcreteCreatorB{}}

	for _, creator := range creators {
		product := creator.CreateProduct()
		product.Use()
	}
}
