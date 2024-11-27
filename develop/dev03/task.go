package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func compare(a, b string, numeric bool) bool {
	if numeric {
		an, err := strconv.Atoi(a)
		if err != nil {
			fmt.Println("Incorrect type if sort(numeric): ", err)
		}

		bn, err := strconv.Atoi(b)
		if err != nil {
			fmt.Println("Incorrect type if sort(numeric): ", err)
		}
		return an < bn
	}

	return a < b

}

func isEqual(a, b []string, numeric bool) bool {
	if len(a) != len(b) {
		return false
	}

	if numeric {
		for i := 0; i < len(a); i++ {
			an, err := strconv.Atoi(a[i])
			if err != nil {
				fmt.Println("Incorrect type if sort(numeric): ", err)
			}

			bn, err := strconv.Atoi(b[i])
			if err != nil {
				fmt.Println("Incorrect type if sort(numeric): ", err)
			}

			if an != bn {
				return false
			}
		}
		return true
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func uniqueSort(words [][]string, numeric bool) (res [][]string) {
	for i := 0; i < len(words); i++ {
		if len(res) == 0 || !isEqual(res[len(res)-1], words[i], numeric) {
			res = append(res, words[i])
		}
	}
	return res
}

func main() {

	k := flag.Int("k", 0, "column number")
	n := flag.Bool("n", false, "compare according to string numerical value")
	r := flag.Bool("r", false, "reverse the result of comparisons")
	u := flag.Bool("u", false, "print unique strings")

	//column - attribut for sort
	//numeric types of sort
	//unique, reverse - change data

	path := flag.String("input", "input.txt", "input file path")

	flag.Parse()

	column := *k
	numeric := *n
	reverse := *r
	unique := *u

	file, err := os.Open(*path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		os.Exit(1)
	}

	maxSizeLine := 0
	for _, line := range lines {
		maxSizeLine = max(maxSizeLine, len(strings.Fields(line)))
	}
	if column < 0 || maxSizeLine <= column {
		fmt.Println("Incorreсt argument k (column number)!")
		os.Exit(1)
	}

	fmt.Println(column, maxSizeLine)

	var words [][]string
	for i := 0; i < len(lines); i++ {
		words = append(words, strings.Fields(lines[i]))
	}

	sort.Slice(words, func(i, j int) bool {
		if len(words[i]) <= column {
			return true
		}
		if len(words[j]) <= column {
			return false
		}
		return compare(words[i][column], words[j][column], numeric)
	})

	if reverse {
		for i := 0; i < len(words)/2; i++ {
			words[i], words[len(words)-i-1] = words[len(words)-i-1], words[i]
		}
	}

	if unique {
		words = uniqueSort(words, numeric)
	}

	var result string
	var resultLoc []string
	for i := 0; i < len(words); i++ {
		resultLoc = append(resultLoc, strings.Join(words[i], " "))
	}
	result = strings.Join(resultLoc, "\n")

	output, err := os.OpenFile((*path)[:len(*path)-4]+"_output.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)

	if err != nil {
		fmt.Println("Error with creating file")
		os.Exit(1)
	}
	_, _ = output.WriteString(result)
	output.Close()

}
