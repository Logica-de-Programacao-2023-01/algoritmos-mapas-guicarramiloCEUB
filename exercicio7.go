package main

import (
	"fmt"
	"strings"
)

func countLetters(word string) map[rune]int {
	letterCount := make(map[rune]int)

	for _, char := range word {
		letterCount[char]++
	}

	return letterCount
}

func countLettersInWords(sentence string) map[string]map[rune]int {
	wordCount := make(map[string]map[rune]int)

	words := strings.Fields(sentence)
	for _, word := range words {
		wordCount[word] = countLetters(word)
	}

	return wordCount
}

func main() {
	sentence := "Hello, world! Welcome to OpenAI."
	wordCount := countLettersInWords(sentence)

	for word, letterCount := range wordCount {
		fmt.Printf("%s: %v\n", word, letterCount)
	}
}
