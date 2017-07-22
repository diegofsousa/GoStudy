package main

import (
	"fmt"
)

func main() {
	/*Só existe "for" :/*/

	a, b := 0, 10

	for a < b {
		a+= 1
	}

	fmt.Println(a)

	/*Forma simples like a C*/

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	/*Omitindo argumentos*/

	var k int
	for k = 0; k < 10; {
		k += 1
	}
	
	l := 0
	for ; l < 10; l++ {
		l += 1
	}

	/*Com range (forma tradicional em Go)*/

	numeros := []int{1, 2, 3, 4, 5}

	for i := range numeros {
		numeros[i] *= 2
	}

	fmt.Println(numeros)

}