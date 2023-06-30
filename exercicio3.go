package main

import "fmt"

func main() {
	m := map[int]int{1: 1, 2: 2, 3: 3, 4: 4, 5: 5}
	fmt.Println(somam(m))
}
func somam(m map[int]int) int {
	soma := 0
	for _, valor := range m {
		soma += valor
	}
	return soma
}
