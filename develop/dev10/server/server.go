package main

import (
	"fmt"
	"net"
	"os"
)

var dict = map[string]string{
	"red":    "красный",
	"green":  "зеленый",
	"blue":   "синий",
	"yellow": "желтый",
}

func main() {

	host := "localhost"
	port := "1234"

	listener, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Server is listening...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			conn.Close()
			continue
		}

		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	for {
		input := make([]byte, 4096)
		n, err := conn.Read(input)
		if n == 0 || err != nil {
			fmt.Println("Error with conn or n:", n, " ", err)
			break
		}

		target := string(input[0:n])
		value, ok := dict[target]
		if !ok {
			value = "unknown"
		}

		fmt.Println(value, " - ", target)

		conn.Write([]byte(value + "\n"))
	}
}
