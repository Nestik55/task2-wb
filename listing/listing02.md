Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
```
2
1
```
Функция test() вернет 2, поскольку Deferred functions may read and assign to the returning function’s named return values.
Или, отложенные функции могут читать и присваивать именованные возвращаемые значения возвращающей функции.
Деферы вызываются в порядке LIFO (стек).
Аргументы отложенной функции оцениваются при вычислении оператора defer.