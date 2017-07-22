package main

import (
	"fmt"
)

func main() {
	/*Declarações de arrays em Go*/
	var a[3]int // Em Go, estes valores ganham automaticamente um valor inicial conhecido como zero value.
	numeros := [5]int{1,2,3,4,5}
	primos := [...]int{2,3,4,7,11,13}
	nomes := [2]string{}

	fmt.Println(a, numeros, primos, nomes)

	/*Declarações de arrays multidimensionais em Go*/
	var multiA [2][2]int
	multiA[0][0], multiA[0][1] = 3, 5
	multiA[1][0], multiA[1][1] = 7, -2

	multiB := [2][2]int{{2,3},{-1,6}}

	fmt.Println("Multi A:", multiA)
	fmt.Println("Multi B:", multiB)

	// Arrays em Go tem sua dinamicidade complicada pois estabelecem limites de memória.
	// Na maioria dos casos usam-se "slices".
}