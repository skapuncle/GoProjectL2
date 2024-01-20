package main

/* Паттерн "Команда" (Command) - это поведенческий паттерн проектирования,
который превращает запросы в объекты, позволяя передавать их как аргументы
при вызове методов, ставить запросы в очередь, логировать их,
а также поддерживать отмену операций.

Преимущества паттерна "Команда" включают:

Декомпозиция запроса в отдельные объекты.
Поддержка отмены операций.
Легкое сохранение истории операций.
Поддержка транзакций и выполнения операций в очереди.

Минусы паттерна "Команда" включают:

Увеличение сложности кода из-за добавления новых классов и интерфейсов.
Неудобство при работе с большим количеством различных типов команд.

*/

import "fmt"

// Command представляет интерфейс для команды.
type Command interface {
	Execute()
}

// ConcreteCommandA реализует интерфейс Command.
type ConcreteCommandA struct {
	receiver Receiver
}

func (c *ConcreteCommandA) Execute() {
	c.receiver.DoSomething()
}

// ConcreteCommandB реализует интерфейс Command.
type ConcreteCommandB struct {
	receiver Receiver
}

func (c *ConcreteCommandB) Execute() {
	c.receiver.DoSomethingElse()
}

// Receiver представляет класс, который выполняет некоторое действие.
type Receiver struct{}

func (r *Receiver) DoSomething() {
	fmt.Println("Receiver: doing something...")
}

func (r *Receiver) DoSomethingElse() {
	fmt.Println("Receiver: doing something else...")
}

// Invoker представляет класс, который выполняет команды.
type Invoker struct {
	commands []Command
}

func (i *Invoker) AddCommand(cmd Command) {
	i.commands = append(i.commands, cmd)
}

func (i *Invoker) ExecuteCommands() {
	for _, cmd := range i.commands {
		cmd.Execute()
	}
}

func main() {
	invoker := &Invoker{}
	receiver := &Receiver{}

	concreteCommandA := &ConcreteCommandA{receiver: receiver}
	concreteCommandB := &ConcreteCommandB{receiver: receiver}

	invoker.AddCommand(concreteCommandA)
	invoker.AddCommand(concreteCommandB)

	invoker.ExecuteCommands()
}
