package main

import "fmt"

/*
«Состояние» - применяется в случаях, когда поведение объекта зависит от его состояния,
и поведение должно изменяться динамически в зависимости от этого состояния.
Он позволяет объекту изменять свое поведение при изменении его внутреннего состояния.

Применимость:
	При разработке игры можно использовать состояние для реализации различных поведений
	персонажей в зависимости от их текущего состояния.

Плюсы:
	Позволяет объекту изменять свое поведение в зависимости от состояния.
	Изолирует код, связанный с определенным состоянием, в отдельных классах.
	Облегчает добавление новых состояний без изменения существующего кода.

Минусы:
	Может привести к увеличению количества классов.
*/

// Интерфейс состояния определяет методы, которые будут вызваны контекстом в зависимости от текущего состояния.
type State interface {
	Handle()
}

// Конкретное состояние A реализует методы для обработки событий в состоянии A.
type ConcreteStateA struct{}

func (s *ConcreteStateA) Handle() {
	fmt.Println("Handling state A")
}

// Конкретное состояние B реализует методы для обработки событий в состоянии B.
type ConcreteStateB struct{}

func (s *ConcreteStateB) Handle() {
	fmt.Println("Handling state B")
}

// Контекст использует состояние для обработки событий.
type ContextState struct {
	state State
}

// Устанавливает текущее состояние контекста.
func (c *ContextState) SetState(state State) {
	c.state = state
}

// Обрабатывает событие в зависимости от текущего состояния.
func (c *ContextState) Request() {
	c.state.Handle()
}

func main() {
	// Создаем контекст состояния
	contextState := &ContextState{}

	// Устанавливаем состояние A
	stateA := &ConcreteStateA{}
	contextState.SetState(stateA)
	// Обрабатываем событие для состояния A
	contextState.Request()

	// Устанавливаем состояние B
	stateB := &ConcreteStateB{}
	contextState.SetState(stateB)
	// Обрабатываем событие для состояния B
	contextState.Request()
}
