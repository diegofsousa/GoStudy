package main

import (
	"fmt"
)

//Default
func ImprimirSaudacao() {
	fmt.Println("Olá Go! :)")
}

//Recebendo parâmetros
func ImprimirMeuNome(nome string) {
	fmt.Printf("Meu nome é %s.\n", nome)
}

//Recebendo parâmetros e retornando valores
func ImprimirSoma(a, b int) int {
	return a + b
}

func main() {
	ImprimirSaudacao()
	ImprimirMeuNome("Diego Fernando")
	fmt.Printf("A soma é: %d.\n", ImprimirSoma(67, 33))
}