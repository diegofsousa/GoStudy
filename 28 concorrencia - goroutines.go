package main

import (
	"fmt"
	"time"
)

func imprimir(n int) {
	for i := 0; i < 3; i++ {
		fmt.Printf("%d ", n)
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	/*A palavra chave "go" abre uma nova rotina de execução*/
	go imprimir(2)
	imprimir(3)
	fmt.Printf("\n")
}