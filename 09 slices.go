package main

import (
	"fmt"
)

func main() {
	/*Declaração dos slices*/
	/*Basta deixar o limite indefinido*/
	var a []int
	primos := []int{2,3,5,7,11,13}
	nomes := []string{}

	/*make (tipo, tamanho inicial, capacidade total)*/
	b := make([]int, 10)
	fmt.Println(b, len(b), cap(b))

	c := make([]int, 10, 20)
	fmt.Println(c, len(c), cap(c))
}