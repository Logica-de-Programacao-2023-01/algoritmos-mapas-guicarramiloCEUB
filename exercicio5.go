package main

import "fmt"

func main() {
	s := "guilherme"
	o := make(map[string]int)
	ocorrencia(s, &o)
	fmt.Println(o)
}

func ocorrencia(s string, mapa *map[string]int) {
	for i := 0; i < len(s); i++ {
		(*mapa)[string(s[i])] += 1
	}
}
