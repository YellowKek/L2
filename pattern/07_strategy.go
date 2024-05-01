package main

import "fmt"

/*
«Стратегия» - применяется в случаях, когда у нас есть схожие алгоритмы, которые могут изменяться динамически
или подходить для разных контекстов. Он позволяет выделить алгоритмы в отдельные классы, делая их
заменяемыми и позволяя выбирать нужный алгоритм на лету.

Применимость:
	Веб-приложение для обработки платежей может использовать стратегию для выбора различных способов оплаты
	в зависимости от предпочтений пользователя.

Плюсы:
	Позволяет изолировать алгоритмы от клиентов, которые их используют.
	Облегчает добавление новых алгоритмов без изменения существующего кода.
	Позволяет заменять алгоритмы на лету.

Минусы:
	Может создать много дополнительных классов.
*/

// Интерфейс стратегии определяет метод для выполнения определенного алгоритма.
type Strategy interface {
	Execute()
}

// Конкретная стратегия A реализует алгоритм A.
type ConcreteStrategyA struct{}

func (s *ConcreteStrategyA) Execute() {
	fmt.Println("Executing strategy A")
}

// Конкретная стратегия B реализует алгоритм B.
type ConcreteStrategyB struct{}

func (s *ConcreteStrategyB) Execute() {
	fmt.Println("Executing strategy B")
}

// Контекст использует стратегию для выполнения определенного алгоритма.
type Context struct {
	strategy Strategy
}

// Устанавливает стратегию контекста.
func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

// Выполняет стратегию.
func (c *Context) ExecuteStrategy() {
	c.strategy.Execute()
}

func main() {
	// Создаем контекст
	context := &Context{}

	// Устанавливаем стратегию A
	strategyA := &ConcreteStrategyA{}
	context.SetStrategy(strategyA)
	// Выполняем стратегию A
	context.ExecuteStrategy()

	// Устанавливаем стратегию B
	strategyB := &ConcreteStrategyB{}
	context.SetStrategy(strategyB)
	// Выполняем стратегию B
	context.ExecuteStrategy()
}
