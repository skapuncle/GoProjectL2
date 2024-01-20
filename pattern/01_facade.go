package main

import (
	"fmt"
	"time"
)

/* Паттерн "Фасад" (Facade) - это структурный паттерн проектирования,
который предоставляет простой интерфейс к сложной системе классов, библиотеке или фреймворку.
Этот паттерн может быть особенно полезен, когда вам нужно представить простой или урезанный
интерфейс к сложной подсистеме.
Преимущества паттерна "Фасад" включают:

Изолирует клиентов от компонентов сложной подсистемы.
Минусы паттерна "Фасад" включают:

Фасад рискует стать "божественным объектом", привязанным ко всем классам программы
*/

// ComplexSubsystemA представляет сложную подсистему A
type ComplexSubsystemA struct{}

func (a *ComplexSubsystemA) OperationA() {
	fmt.Println("Operation A")
}

// ComplexSubsystemB представляет сложную подсистему B
type ComplexSubsystemB struct{}

func (b *ComplexSubsystemB) OperationB() {
	fmt.Println("Operation B")
}

// Facade представляет фасад, который скрывает сложность подсистем
type Facade struct {
	a *ComplexSubsystemA
	b *ComplexSubsystemB
}

func NewFacade() *Facade {
	return &Facade{
		a: &ComplexSubsystemA{},
		b: &ComplexSubsystemB{},
	}
}

func (f *Facade) Operation() {
	f.a.OperationA()
	time.Sleep(1 * time.Second)
	f.b.OperationB()
}

func main() {
	f := NewFacade()
	f.Operation()
}
