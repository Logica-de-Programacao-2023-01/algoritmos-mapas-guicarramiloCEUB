package main

import (
	"fmt"
	"strings"
)

func countWords(input string) map[string]int {
	wordCount := make(map[string]int)

	words := strings.Fields(input)

	for _, word := range words {
		wordCount[word]++
	}

	return wordCount
}

func main() {
	input := "Esta é uma frase de teste. É uma frase simples, mas uma frase significativa."

	wordCount := countWords(input)

	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}
}

