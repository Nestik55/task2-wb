package main

import (
	"errors"
	"strconv"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func UnpackString(str []rune) ([]rune, error) {
	var res []rune

	var printSlash bool = false
	var prevIsSl bool = false
	var prev rune = -1
	for _, cur := range str {
		if prev == -1 {
			prev = cur
			continue
		}

		if unicode.IsDigit(cur) {
			digit, err := strconv.Atoi(string(cur))
			if err != nil {
				return []rune{}, err
			}

			if prev == '\\' {
				if printSlash {
					for ; digit > 0; digit-- {
						res = append(res, prev)
					}
					printSlash = false
					prev = -1
					continue
				}
				prev = cur
				prevIsSl = true
			} else if unicode.IsDigit(prev) && prevIsSl {
				for ; digit > 0; digit-- {
					res = append(res, prev)
				}
				prev = -1
				prevIsSl = false
			} else if !unicode.IsDigit(prev) {
				for ; digit > 0; digit-- {
					res = append(res, prev)
				}
				prev = -1
				prevIsSl = false
			} else {
				return []rune{}, errors.New("Incorrect string")
			}
		} else if cur == '\\' {
			if prev == '\\' {
				printSlash = true
				continue
			}
			res = append(res, prev)
			prev = cur
		} else {
			if prev != -1 {
				res = append(res, prev)
			}
			prev = cur
		}
	}

	if prev != -1 {
		res = append(res, prev)
	}
	if prev == '\\' {
		return []rune{}, errors.New("Incorrect string")
	}

	return res, nil
}
