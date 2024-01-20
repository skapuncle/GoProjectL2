package main

/* Паттерн "Цепочка вызовов" (Chain of Responsibility) - это поведенческий паттерн
проектирования, который позволяет передавать запросы последовательно по цепочке
обработчиков. Каждый последующий обработчик решает, может ли он обработать запрос
сам и стоит ли передавать запрос дальше по цепи.

Преимущества паттерна "Цепочка обязанностей" включают:

- Избавление от жёсткой привязки отправителя запроса к его получателю, позволяя выстраивать цепь из различных обработчиков динамически.
- Возможность передачи запроса последовательно по цепочке обработчиков.

Минусы паттерна "Цепочка обязанностей" включают:

- Увеличение сложности кода из-за добавления новых классов и интерфейсов.
- Неудобство при работе с большим количеством различных типов обработчиков.
*/

import (
	"fmt"
)

// Handler представляет интерфейс для обработчика.
type Handler interface {
	SetNext(Handler) Handler
	HandleRequest(request string) string
}

// BaseHandler реализует интерфейс Handler.
type BaseHandler struct {
	next Handler
}

func (b *BaseHandler) SetNext(handler Handler) Handler {
	b.next = handler
	return handler
}

func (b *BaseHandler) HandleRequest(request string) string {
	if b.next != nil {
		return b.next.HandleRequest(request)
	}
	return ""
}

// ConcreteHandlerA реализует интерфейс Handler.
type ConcreteHandlerA struct {
	BaseHandler
}

func (c *ConcreteHandlerA) HandleRequest(request string) string {
	if request == "request A" {
		return fmt.Sprintf("Handled by %s", "ConcreteHandlerA")
	}
	return c.BaseHandler.HandleRequest(request)
}

// ConcreteHandlerB реализует интерфейс Handler.
type ConcreteHandlerB struct {
	BaseHandler
}

func (c *ConcreteHandlerB) HandleRequest(request string) string {
	if request == "request B" {
		return fmt.Sprintf("Handled by %s", "ConcreteHandlerB")
	}
	return c.BaseHandler.HandleRequest(request)
}

func main() {
	handlerA := &ConcreteHandlerA{}
	handlerB := &ConcreteHandlerB{}

	handlerA.SetNext(handlerB)

	requests := []string{"request A", "request B", "request C"}
	for _, request := range requests {
		fmt.Printf("Request: %s\n%s\n", request, handlerA.HandleRequest(request))
	}
}
