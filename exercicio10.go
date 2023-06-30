package main

import "fmt"

func countNumberPairs(numbers []int) map[[2]int]int {
	numberPairs := make(map[[2]int]int)

	for i := 0; i < len(numbers)-1; i++ {
		for j := i + 1; j < len(numbers); j++ {
			pair := [2]int{numbers[i], numbers[j]}
			numberPairs[pair]++
		}
	}

	return numberPairs
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 1, 2, 1}
	numberPairCounts := countNumberPairs(numbers)

	for pair, count := range numberPairCounts {
		fmt.Printf("%v: %d\n", pair, count)
	}
}
