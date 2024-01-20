package main

/* Паттерн "Стратегия" (Strategy) - это поведенческий паттерн проектирования,
который определяет семейство схожих алгоритмов и помещает каждый из них в
собственный класс, после чего алгоритмы можно взаимозаменять прямо во время
исполнения программы.

Преимущества паттерна "Стратегия" включают:

Горячую замену алгоритмов на лету.
Изоляцию кода и данных алгоритмов от остальных классов.

Минусы паттерна "Стратегия" включают:

Усложнение программы за счет дополнительных классов.
Необходимость для клиента знать, в чём состоит разница между стратегиями, чтобы выбрать подходящую.
*/

import "fmt"

// Strategy представляет интерфейс для стратегии.
type Strategy interface {
	DoAlgorithm()
}

// ConcreteStrategyA реализует интерфейс Strategy.
type ConcreteStrategyA struct{}

func (cs *ConcreteStrategyA) DoAlgorithm() {
	fmt.Println("ConcreteStrategyA: executing algorithm")
}

// ConcreteStrategyB реализует интерфейс Strategy.
type ConcreteStrategyB struct{}

func (cs *ConcreteStrategyB) DoAlgorithm() {
	fmt.Println("ConcreteStrategyB: executing algorithm")
}

// Context представляет контекст, который использует стратегию.
type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) DoSomeBusinessLogic() {
	c.strategy.DoAlgorithm()
}

func main() {
	context := &Context{}

	strategyA := &ConcreteStrategyA{}
	context.SetStrategy(strategyA)
	context.DoSomeBusinessLogic()

	strategyB := &ConcreteStrategyB{}
	context.SetStrategy(strategyB)
	context.DoSomeBusinessLogic()
}
