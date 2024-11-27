package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port
go-telnet mysite.ru 8080
go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу {через аргумент --timeout, по умолчанию 10s}.

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

func main() {
	timeout := flag.Duration("timeout", 10*time.Second, "use timeout")
	flag.Parse()

	args := flag.Args()

	if len(args) != 2 {
		fmt.Println("You must use this: [host] [port]")
		os.Exit(1)
	}

	host := args[0]
	port := args[1]

	conn, err := net.DialTimeout("tcp", host+":"+port, *timeout)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close()

	go handleClose(conn)

	for {
		var source string
		fmt.Print("<<<")
		_, err := fmt.Scan(&source)

		if errors.Is(err, io.EOF) {
			fmt.Println("you exit")
			os.Exit(0)
		}
		if err != nil {
			fmt.Println(err)
			continue
		}

		if _, err := conn.Write([]byte(source)); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Print(">>>")

		buff := make([]byte, 1024)
		n, err := conn.Read(buff)
		if err != nil {
			break
		}

		fmt.Println(string(buff[0:n]))
	}

}

func handleClose(conn net.Conn) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)

	<-signalChan
	func() {
		_ = conn.Close()
	}()

	os.Exit(0)
}
