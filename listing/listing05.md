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
```
error
```

Интерфейс равен nil, только если и тип, и значение равны nil. В данном случае функция test возвращает интерфейс, в котором данные будут nil, а тип будет определен и не будет равняться nil.
