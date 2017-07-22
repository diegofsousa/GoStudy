package main

import "fmt"

func imprimirDados(nome string, idade int) {
	fmt.Printf("%s tem %d anos.\n", nome, idade)
}

func main() {
	imprimirDados("Diego Fernando", 21)
}