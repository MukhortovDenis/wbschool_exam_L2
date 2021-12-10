package main

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

func searchAnagramma(dict *[]string) *map[string][]string {
	dictMap := make(map[string][]string)
	runeVal := make(map[string]int)
	runeBool := make(map[string]bool)
	for _, word := range *dict {
		word = strings.ToLower(word)
		atRune := []rune(word)
		var val int
		for i := range atRune {
			val += int(atRune[i])
		}
		runeVal[word] = val

	}
	for _, elem := range *dict {
		elem = strings.ToLower(elem)
		slice := []string{}
		if !runeBool[elem] {
			slice = append(slice, elem)
			runeBool[elem] = true
		}
		for key2, value2 := range runeVal {
			if elem != key2 && runeVal[elem] == value2 && len(elem) == len(key2) && !runeBool[key2] {
				slice = append(slice, key2)
				runeBool[key2] = true
			}
		}
		if len(slice) > 1 {
			sort.Strings(slice)
			dictMap[elem] = slice
		}
	}
	return &dictMap
}
