package main

import "fmt"

/*
«Цепочка вызовов» - используется, когда имеется более одного объекта, который может обработать определенный запрос.

Применимость:
	Веб-приложение может использовать цепочку обработчиков запросов для авторизации,
	валидации и обработки запросов на различных уровнях приложения.

Плюсы:
	Позволяет гибко настраивать порядок обработки запросов.
	Изолирует отправителя запроса от его получателя.

Минусы:
	Не гарантирует, что запрос будет обработан.
	Может стать сложным для отладки.
*/

// Handler определяет интерфейс обработчика запроса.
type Handler interface {
	SetNext(handler Handler)
	HandleRequest(request int)
}

// ConcreteHandler представляет конкретную реализацию обработчика.
type ConcreteHandler struct {
	next Handler
}

func (c *ConcreteHandler) SetNext(handler Handler) {
	c.next = handler
}

func (c *ConcreteHandler) HandleRequest(request int) {
	fmt.Printf("ConcreteHandler: Handling request %d\n", request)
	if c.next != nil {
		c.next.HandleRequest(request)
	}
}

func main() {
	// Создаем цепочку обработчиков.
	handler1 := &ConcreteHandler{}
	handler2 := &ConcreteHandler{}
	handler3 := &ConcreteHandler{}

	handler1.SetNext(handler2)
	handler2.SetNext(handler3)

	// Посылаем запрос по цепочке обработчиков.
	handler1.HandleRequest(1)
}
