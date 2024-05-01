package main

import "fmt"

/* «Фасад» - структурный паттерн проектирования, который предоставляет простой интерфейс к сложной подсистеме,
облегчая использование этой подсистемы. Фасад выступает в качестве посредника между клиентом и различными
компонентами подсистемы, скрывая сложность и детали реализации от клиента.

Используется:
	где необходим простой интерфейс к сложной системе, состоящей из множества взаимосвязанных объектов
	работа с устаревшими системами или сторонними библиотеками

Плюсы:
	Упрощает работу с комплексными системами, скрывая их сложность от клиента.
	Снижает связность между компонентами системы, что упрощает поддержку и изменение кода.
	Повышает переиспользуемость кода.

Минусы:
	Может привести к увеличению количества классов в системе, если фасады необходимы для разных сценариев использования.
	Неправильное применение может привести к созданию излишне сложных фасадов.
*/

// SubsystemA представляет компонент A подсистемы.
type SubsystemA struct{}

func (s *SubsystemA) OperationA() {
	fmt.Println("Subsystem A: Operation A")
}

// SubsystemB представляет компонент B подсистемы.
type SubsystemB struct{}

func (s *SubsystemB) OperationB() {
	fmt.Println("Subsystem B: Operation B")
}

// Facade представляет собой фасад, скрывающий сложность подсистемы.
type Facade struct {
	subsystemA *SubsystemA
	subsystemB *SubsystemB
}

func NewFacade() *Facade {
	return &Facade{
		subsystemA: &SubsystemA{},
		subsystemB: &SubsystemB{},
	}
}

// Operation делегирует запросы клиента соответствующим объектам внутри подсистемы.
func (f *Facade) Operation() {
	fmt.Println("Facade: Initialization and orchestration of subsystems")
	f.subsystemA.OperationA()
	f.subsystemB.OperationB()
}

func main() {
	facade := NewFacade()
	facade.Operation()
}
