package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	//A, B, C, c, n - for output
	//i - change data fot search
	//v - invert result
	//F - fix search
	var after, before, context int
	var count, ignore, inverse, full, lineNum bool

	flag.IntVar(&after, "A", 0, "output n line after")
	flag.IntVar(&before, "B", 0, "output n line before")
	flag.IntVar(&context, "C", 0, "output context (n before and after)")
	flag.BoolVar(&count, "c", false, "output count equal line") //
	flag.BoolVar(&ignore, "i", false, "ignore register")        //
	flag.BoolVar(&inverse, "v", false, "output inverse")        //
	flag.BoolVar(&full, "F", false, "full match")               //
	flag.BoolVar(&lineNum, "n", false, "output number line with equal")
	flag.Parse()

	pattern := flag.Arg(0)
	if ignore {
		pattern = `(?i)` + pattern
	}

	if full {
		pattern = `^` + pattern + `$`
	}

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
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

	_, err = regexp.Compile(pattern)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cnt := 0
	var numberResLines []int
	for i, line := range lines {
		ok, err := regexp.MatchString(pattern, line)
		if err != nil {
			fmt.Printf("Error with regexp: pattern - %v, line - %v\n", pattern, line)
			continue
		}

		if (inverse && !ok) || (!inverse && ok) {
			numberResLines = append(numberResLines, i)
			cnt++
		}
	}

	if count {
		fmt.Printf("count of this lines:%v\n\n", cnt)
	}

	for i := 0; i < len(numberResLines); i++ {
		if context > 0 {
			before = context
			after = context
		}
		printBefore(numberResLines[i], before, lines, lineNum)

		if lineNum {
			fmt.Print(i+1, " ")
		}
		fmt.Println(lines[numberResLines[i]])

		printAfter(numberResLines[i], after, lines, lineNum)
		fmt.Println()
	}

}

func printBefore(index, beforeN int, lines []string, n bool) {

	for i := max(0, index-beforeN); i < index; i++ {
		if n {
			fmt.Print(i+1, " ")
		}
		fmt.Println(lines[i])
	}
}

func printAfter(index, afterN int, lines []string, n bool) {

	for i := index + 1; i < min(index+afterN+1, len(lines)); i++ {
		if n {
			fmt.Print(i+1, " ")
		}
		fmt.Println(lines[i])
	}
}
