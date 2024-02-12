package main

import (
	"fmt"
	"sort"
	"strings"
)

func FindAnagramSets(dict *[]string) *map[string][]string {
	anagramSets := make(map[string][]string)

	for _, word := range *dict {
		// Приведение слова к нижнему регистру
		word = strings.ToLower(word)

		// Сортировка букв в слове
		sortedWord := sortString(word)

		// Добавление слова в соответствующее множество анаграмм
		anagramSets[sortedWord] = append(anagramSets[sortedWord], word)
	}

	// Удаление множества из одного элемента из результатов
	for key, value := range anagramSets {
		if len(value) <= 1 {
			delete(anagramSets, key)
		} else {
			// Сортировка массива слов в множестве по возрастанию
			sort.Strings(anagramSets[key])
		}
	}

	return &anagramSets
}

// Вспомогательная функция для сортировки букв в слове
func sortString(s string) string {
	sChars := strings.Split(s, "")
	sort.Strings(sChars)
	return strings.Join(sChars, "")
}

func main() {
	dictionary := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	anagramSets := FindAnagramSets(&dictionary)

	for _, value := range *anagramSets {
		fmt.Printf("Множество анаграмм для слова %s: %v\n", value[0], value)
	}
}
