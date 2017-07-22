package main

import (
	"fmt"
)

func main() {
	s := make([]int, 0)

	//inserção no fim
	s = append(s, 23)
	fmt.Println(s)

	//inserção no início
	novoElemento := []int{22}
	s = append(novoElemento, s...)
	fmt.Println(s)

	//remoção no início
	j := []int{20, 30, 40, 50, 60}
	j = j[1:]
	fmt.Println(j)

	//remoção no fim
	k := []int{20, 30, 40, 50, 60}
	k = k[:3]
	fmt.Println(k)

	/*
		criando uma cópia do slice
		escopo: func copy(destino, origem []Tipo) int
	*/
	numeros := []int{1, 2, 3, 4, 5}
	dobros :=make([]int, len(numeros))

	copy(dobros, numeros)

	for i := range dobros{
		dobros[i] *= 2
	}

	fmt.Println(numeros)
	fmt.Println(dobros)
}