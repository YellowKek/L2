package main

import "fmt"

/*
«Команда» - используется для инкапсуляции запроса в виде объекта, что позволяет параметризовать клиентов
с организацией очереди запросов, регистрацией запросов и поддержкой отмены операций.

Применимость:
	В текстовом редакторе можно реализовать команды для выполнения операций над текстом:
	создание, удаление, копирование, вставка и так далее.

Плюсы:
	Изоляция отправителя и получателя.
	Поддержка отмены операций.
	Поддержка выполнения команды в другом потоке.

Минусы:
	Может привести к увеличению числа классов.
*/

// Command определяет интерфейс команды.
type Command interface {
	Execute()
}

// ConcreteCommand представляет конкретную реализацию команды.
type ConcreteCommand struct {
	receiver *Receiver
}

func NewConcreteCommand(receiver *Receiver) *ConcreteCommand {
	return &ConcreteCommand{
		receiver: receiver,
	}
}

func (c *ConcreteCommand) Execute() {
	c.receiver.Action()
}

// Receiver представляет получателя команды.
type Receiver struct{}

func (r *Receiver) Action() {
	fmt.Println("Receiver: Performing action")
}

// Invoker представляет вызывающего, который выполняет команду.
type Invoker struct {
	command Command
}

func (i *Invoker) SetCommand(command Command) {
	i.command = command
}

func (i *Invoker) ExecuteCommand() {
	fmt.Println("Invoker: Execute command")
	i.command.Execute()
}

func main() {
	receiver := &Receiver{}
	command := NewConcreteCommand(receiver)

	// Создаем вызывающего и передаем ему команду.
	invoker := &Invoker{}
	invoker.SetCommand(command)

	// Вызываем команду через вызывающего.
	invoker.ExecuteCommand()
}
