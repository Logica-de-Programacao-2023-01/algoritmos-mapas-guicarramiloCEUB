package main

import "fmt"

func generateFibonacci(n int) map[int]int {
	fibonacci := make(map[int]int)

	fibonacci[0] = 0
	if n > 0 {
		fibonacci[1] = 1
	}

	for i := 2; fibonacci[i-1] <= n; i++ {
		fibonacci[i] = fibonacci[i-1] + fibonacci[i-2]
	}

	return fibonacci
}

func main() {
	n := 100
	fibonacciSequence := generateFibonacci(n)

	for index, number := range fibonacciSequence {
		fmt.Printf("Index: %d, Number: %d\n", index, number)
	}
}
