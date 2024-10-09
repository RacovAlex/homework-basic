package main

import (
	"strings"
	"unicode"
)

func main() {
	// Place your code here.
}

func normalizeWord(word string) string {
	// Приводим к нижнему регистру.
	word = strings.ToLower(word)
	// Удаляем знаки препинания с краев слов.
	word = strings.TrimFunc(word, unicode.IsPunct)
	return word
}

func countWords(str string) map[string]int {
	if len(str) == 0 {
		return map[string]int{}
	}

	counts := make(map[string]int)
	for _, word := range strings.Fields(str) {
		word = normalizeWord(word)
		if word != "" {
			counts[word]++
		}
	}
	return counts
}
