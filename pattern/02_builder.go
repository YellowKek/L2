package main

import "fmt"

/*
«Строитель» — это порождающий паттерн проектирования, который позволяет создавать сложные объекты пошагово.
Он предоставляет способ создания объекта с различными конфигурациями, не увеличивая сложность конструктора.

Применимость:
	когда процесс создания объекта состоит из множества шагов, которые могут варьироваться в зависимости от нужд клиента.
	когда требуется создавать различные представления одного и того же объекта.

Плюсы:
	Упрощает создание сложных объектов, скрывая детали их конструирования.
	Позволяет создавать различные конфигурации объектов, не усложняя их классы.
	Снижает зависимость между строителем и конечным продуктом.

Минусы:
	Может быть избыточным для создания простых объектов.
	Увеличивает количество классов в системе.
*/

// Product представляет продукт, который мы строим.
type ProductB struct {
	PartA string
	PartB string
}

// Builder определяет интерфейс для создания продукта.
type Builder interface {
	BuildPartA()
	BuildPartB()
	GetProduct() *ProductB
}

// ConcreteBuilder представляет конкретную реализацию строителя.
type ConcreteBuilder struct {
	product *ProductB
}

func NewConcreteBuilder() *ConcreteBuilder {
	return &ConcreteBuilder{
		product: &ProductB{},
	}
}

func (b *ConcreteBuilder) BuildPartA() {
	b.product.PartA = "Part A"
}

func (b *ConcreteBuilder) BuildPartB() {
	b.product.PartB = "Part B"
}

func (b *ConcreteBuilder) GetProduct() *ProductB {
	return b.product
}

// Director отвечает за конструирование объекта, используя строитель.
type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{
		builder: builder,
	}
}

func (d *Director) Construct() *ProductB {
	d.builder.BuildPartA()
	d.builder.BuildPartB()
	return d.builder.GetProduct()
}

func main() {
	// Создание строителя и директора.
	builder := NewConcreteBuilder()
	director := NewDirector(builder)

	// Строим продукт с помощью директора.
	product := director.Construct()

	fmt.Println("Product Parts:", product.PartA, product.PartB)
}
