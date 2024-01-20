package main

/*Паттерн "Посетитель" (Visitor) - это поведенческий паттерн проектирования,
который позволяет добавлять новые операции в программу, не изменяя классы объектов,
над которыми эти операции могут выполняться.

Преимущества паттерна "Посетитель" включают:

Добавление новых операций без изменения существующих классов.
Предоставление возможности для разных алгоритмов обработки объектов.
Минусы паттерна "Посетитель" включают:

Увеличение сложности кода из-за добавления новых классов и интерфейсов.
Неудобство при работе с большим количеством различных типов объектов.
*/

import "fmt"

// Visitor представляет интерфейс для посетителя.
type Visitor interface {
	VisitConcreteElementA(*ConcreteElementA)
	VisitConcreteElementB(*ConcreteElementB)
}

// ConcreteVisitorA реализует интерфейс Visitor.
type ConcreteVisitorA struct{}

func (v *ConcreteVisitorA) VisitConcreteElementA(element *ConcreteElementA) {
	fmt.Printf("ConcreteVisitorA: visiting %s\n", element.Name())
}

func (v *ConcreteVisitorA) VisitConcreteElementB(element *ConcreteElementB) {
	fmt.Printf("ConcreteVisitorA: visiting %s\n", element.Name())
}

// ConcreteVisitorB реализует интерфейс Visitor.
type ConcreteVisitorB struct{}

func (v *ConcreteVisitorB) VisitConcreteElementA(element *ConcreteElementA) {
	fmt.Printf("ConcreteVisitorB: visiting %s\n", element.Name())
}

func (v *ConcreteVisitorB) VisitConcreteElementB(element *ConcreteElementB) {
	fmt.Printf("ConcreteVisitorB: visiting %s\n", element.Name())
}

// Element представляет интерфейс для элемента.
type Element interface {
	Accept(visitor Visitor)
}

// ConcreteElementA реализует интерфейс Element.
type ConcreteElementA struct{}

func (e *ConcreteElementA) Accept(visitor Visitor) {
	visitor.VisitConcreteElementA(e)
}

func (e *ConcreteElementA) Name() string {
	return "ConcreteElementA"
}

// ConcreteElementB реализует интерфейс Element.
type ConcreteElementB struct{}

func (e *ConcreteElementB) Accept(visitor Visitor) {
	visitor.VisitConcreteElementB(e)
}

func (e *ConcreteElementB) Name() string {
	return "ConcreteElementB"
}

func main() {
	elements := []Element{&ConcreteElementA{}, &ConcreteElementB{}}
	visitors := []Visitor{&ConcreteVisitorA{}, &ConcreteVisitorB{}}

	for _, element := range elements {
		for _, visitor := range visitors {
			element.Accept(visitor)
		}
	}
}
