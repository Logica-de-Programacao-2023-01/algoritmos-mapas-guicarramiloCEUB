package main

import "fmt"

func main() {
	m1 := map[string]int{
		"gui":  20,
		"bia":  22,
		"mika": 11,
		"fe":   17,
	}
	m2 := map[string]int{
		"gui":    20,
		"bia":    22,
		"mika":   11,
		"fe":     17,
		"fred":   11,
		"raquel": 23,
	}
	junto := juntar(&m1, &m2)
	fmt.Println(junto)
}

func juntar(m1 *map[string]int, m2 *map[string]int) map[string]int {
	junto := make(map[string]int)
	for chave, valor := range *m1 {
		junto[chave] = valor
	}
	for chave, valor := range *m2 {
		junto[chave] = valor
	}
	return junto
}
