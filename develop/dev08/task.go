package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func doCommand(command string) bool {
	commands := strings.Split(command, "|")

	for _, com := range commands {
		if len(com) == 0 {
			continue
		}
		words := strings.Fields(com)
		switch words[0] {
		case "quit":
			return true
		case "cd":
			if len(words) < 2 {
				fmt.Println("cd: need argument(directory)\n")
				return true
			}
			err := os.Chdir(words[1])
			if err != nil {
				fmt.Println("cd: ", err)
				return true
			}

		case "pwd":
			dir, err := os.Getwd()
			if err != nil {
				fmt.Println("pwd: ", err)
				return true
			}
			fmt.Println(dir)
		case "echo":
			if len(words) < 2 {
				fmt.Println("echo: need argument")
				return true
			}
			for i := 1; i < len(words); i++ {
				fmt.Print(words[1], " ")
			}
			fmt.Println()
		case "kill":
			if len(words) < 2 {
				fmt.Println("kill: need argument")
				return true
			}
			out, err := exec.Command("kill", words[1]).Output()
			if err != nil {
				fmt.Println("kill: ", err)
				return true
			}
			fmt.Println(string(out))
		case "ps":
			out, err := exec.Command("ps").Output()
			if err != nil {
				fmt.Println("ps: ", err)
				return true
			}
			fmt.Println(string(out))
		default:
			fmt.Println("Incorrect command: ", words)
		}
	}

	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(">")
		scanner.Scan()

		if doCommand(scanner.Text()) {
			break
		}
	}
}
