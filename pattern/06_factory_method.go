package main

import "fmt"

/*
«Фабричный метод» - применяется, когда у нас есть суперкласс с несколькими подклассами и
в зависимости от условий мы хотим вернуть один из подклассов.

Применимость:
	При разработке игры с различными типами персонажей,
	каждый тип персонажа может иметь свой фабричный метод для создания соответствующего объекта.

Плюсы:
	Изолирует создание объектов от использования.
	Упрощает добавление новых продуктов.
	Позволяет гарантировать соответствие создаваемых объектов интерфейсу.

Минусы:
	Может привести к большому количеству подклассов, если классы-продукты различаются только в деталях.
*/

// ProductF определяет интерфейс продукта.
type ProductF interface {
	Use()
}

// ConcreteProduct представляет конкретную реализацию продукта.
type ConcreteProduct struct{}

func (p *ConcreteProduct) Use() {
	fmt.Println("Using concrete product")
}

// Creator определяет интерфейс создателя.
type Creator interface {
	CreateProduct() ProductF
}

// ConcreteCreator представляет конкретную реализацию создателя.
type ConcreteCreator struct{}

func (c *ConcreteCreator) CreateProduct() ProductF {
	fmt.Println("Creating concrete product")
	return &ConcreteProduct{}
}

func main() {
	// Создаем создателя и используем его для создания продукта.
	creator := &ConcreteCreator{}
	product := creator.CreateProduct()
	product.Use()
}
