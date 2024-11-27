package task

import (
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func task(words *[]string) map[string][]string {
	mp := make(map[string][]string)

	for _, word := range *words {
		runedWord := []rune(strings.ToLower(word))

		if len(runedWord) <= 1 {
			continue
		}

		sort.Slice(runedWord, func(i, j int) bool {
			return runedWord[i] < runedWord[j]
		})

		sortedWord := string(runedWord)

		mp[sortedWord] = append(mp[sortedWord], word)
	}

	res := make(map[string][]string)
	for k := range mp {
		var forSort []string
		newKey := mp[k][0]
		for j := 0; j < len(mp[k]); j++ {
			forSort = append(forSort, strings.ToLower(mp[k][j]))
		}

		sort.Slice(forSort, func(i, j int) bool {
			return forSort[i] < forSort[j]
		})

		for j := 0; j < len(forSort); j++ {
			if len(res[newKey]) == 0 || res[newKey][len(res[newKey])-1] != forSort[j] {
				res[newKey] = append(res[newKey], forSort[j])
			}
		}
	}

	return res
}
