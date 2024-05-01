package main

import "fmt"

/*
«Посетитель» — это поведенческий паттерн проектирования, который позволяет добавлять новые операции к объектам
без изменения их классов. Посетитель позволяет определить новую операцию, не меняя классы объектов,
к которым она применяется.

Применимость:
	когда необходимо добавить новые операции к объектам, но изменение классов объектов нежелательно или невозможно.
	когда классы объектов часто изменяются, но операции, которые выполняются над ними, стабильны.

Плюсы:
	Позволяет добавлять новые операции к объектам без изменения их классов.
	Способствует разделению обязанностей между классами.
	Упрощает добавление новой функциональности к объектам.

Минусы:
	Может привести к увеличению количества классов в системе.
	Усложняет структуру кода из-за необходимости создания посетителей для каждого нового типа объекта.
*/

// Element представляет интерфейс элемента, который может быть посещен.
type Element interface {
	Accept(visitor Visitor)
}

// ConcreteElementA и ConcreteElementB представляют конкретные элементы.
type ConcreteElementA struct{}

func (e *ConcreteElementA) Accept(visitor Visitor) {
	visitor.VisitConcreteElementA(e)
}

type ConcreteElementB struct{}

func (e *ConcreteElementB) Accept(visitor Visitor) {
	visitor.VisitConcreteElementB(e)
}

// Visitor определяет интерфейс посетителя.
type Visitor interface {
	VisitConcreteElementA(element *ConcreteElementA)
	VisitConcreteElementB(element *ConcreteElementB)
}

// ConcreteVisitor представляет конкретного посетителя.
type ConcreteVisitor struct{}

func (v *ConcreteVisitor) VisitConcreteElementA(element *ConcreteElementA) {
	fmt.Println("Visitor visits ConcreteElementA")
}

func (v *ConcreteVisitor) VisitConcreteElementB(element *ConcreteElementB) {
	fmt.Println("Visitor visits ConcreteElementB")
}

// ObjectStructure представляет структуру объектов, которые могут быть посещены.
type ObjectStructure struct {
	elements []Element
}

func (o *ObjectStructure) Attach(element Element) {
	o.elements = append(o.elements, element)
}

func (o *ObjectStructure) Accept(visitor Visitor) {
	for _, element := range o.elements {
		element.Accept(visitor)
	}
}

func main() {
	elementA := &ConcreteElementA{}
	elementB := &ConcreteElementB{}
	objectStructure := &ObjectStructure{}
	objectStructure.Attach(elementA)
	objectStructure.Attach(elementB)

	// Посещение объектов с помощью посетителя.
	visitor := &ConcreteVisitor{}
	objectStructure.Accept(visitor)
}
