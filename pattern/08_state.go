package main

/* Паттерн "Состояние" (State) - это поведенческий паттерн проектирования,
который позволяет объектам менять поведение в зависимости от своего состояния.

Преимущества паттерна "Состояние" включают:

- Изменение поведения объекта в зависимости от его состояния.
- Упрощение добавления новых состояний и модификации поведения объекта.

Минусы паттерна "Состояние" включают:

- Увеличение сложности кода из-за добавления новых классов и интерфейсов.
- Неудобство при работе с большим количеством различных типов состояний.
*/

import "fmt"

// State представляет интерфейс для состояния.
type State interface {
	Handle()
}

// ConcreteStateA реализует интерфейс State.
type ConcreteStateA struct{}

func (cs *ConcreteStateA) Handle() {
	fmt.Println("Handling state A")
}

// ConcreteStateB реализует интерфейс State.
type ConcreteStateB struct{}

func (cs *ConcreteStateB) Handle() {
	fmt.Println("Handling state B")
}

// Context представляет контекст, который использует состояние.
type Context struct {
	state State
}

func (c *Context) SetState(state State) {
	c.state = state
}

func (c *Context) Request() {
	c.state.Handle()
}

func main() {
	context := &Context{}

	stateA := &ConcreteStateA{}
	context.SetState(stateA)
	context.Request()

	stateB := &ConcreteStateB{}
	context.SetState(stateB)
	context.Request()
}
