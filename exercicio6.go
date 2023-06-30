package main

import "fmt"

func mergeWordCounts(wordCountsList []map[string]int) map[string]int {
	mergedWordCount := make(map[string]int)

	for _, wordCounts := range wordCountsList {
		for word, count := range wordCounts {
			mergedWordCount[word] += count
		}
	}

	return mergedWordCount
}

func main() {
	wordCountsList := []map[string]int{
		{"apple": 2, "banana": 3, "orange": 1},
		{"apple": 1, "banana": 2, "grape": 4},
		{"banana": 5, "grape": 3, "kiwi": 2},
	}

	mergedCount := mergeWordCounts(wordCountsList)
	fmt.Println(mergedCount)
}
