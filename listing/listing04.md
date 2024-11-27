Что выведет программа? Объяснить вывод программы.

```go
package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for n := range ch {
		println(n)
	}
}
```

Ответ:
```
Программа выведет числа от 0 до 9, затем словит дедлок, поскольку в цикле бесконечно ожидаются данные, которые туда уже никогда не поступят, чтобы это исправить, необходимо закрыть канал в горутине после цикла.

```