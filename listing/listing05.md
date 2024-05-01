Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}

```

Ответ:

Программа выведет error.

Так как функция test() возвращает тип *customError, даже если она возвращает nil, тип ошибки не является nil. Это не
встроенный тип данных, такой как error, который возвращает nil, когда ошибки нет.

Поэтому, даже если test() возвращает nil, переменная err всё равно содержит значение типа *customError, которое не равно
nil. И поскольку в условии проверяется наличие ошибки, программа выведет "error".
