package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	var delimiter string
	var fields string
	var separated bool

	flag.StringVar(&delimiter, "d", "\t", "choose delimiter")
	flag.StringVar(&fields, "f", "", "choose fields")
	flag.BoolVar(&separated, "s", false, "do not print lines that do not contain the field separator character")
	flag.Parse()
	fmt.Println(fields, separated, delimiter)
	scanner := bufio.NewScanner(os.Stdin)

	var words [][]string
	maxSize := 0
	for scanner.Scan() {
		text := strings.Split(scanner.Text(), delimiter)
		maxSize = max(maxSize, len(text))

		words = append(words, text)
	}

	if fields == "" {
		for i := 0; i < len(words); i++ {
			if separated && len(words[i]) <= 1 {
				continue
			}
			for _, word := range words[i] {
				fmt.Print(word + "\t")
			}
			fmt.Println()
		}
		return
	}

	intervals := strings.Split(fields, ",")
	var columns []struct {
		l int
		r int
	}

	for i := range intervals {
		interval := strings.Split(intervals[i], "-")

		left, err := strconv.Atoi(interval[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		right, err := strconv.Atoi(interval[len(interval)-1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		columns = append(columns, struct {
			l int
			r int
		}{l: left - 1, r: right - 1})
	}

	for _, word := range words {
		for _, column := range columns {
			for i := column.l; i <= column.r; i++ {
				if i < len(word) {
					fmt.Print(word[i] + "\t")
				}
			}
		}
		fmt.Println()
	}

}
