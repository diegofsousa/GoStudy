package main

import (
	"fmt"
)

func main() {
	//Criamos um canal com tamanho do buffer igual a 3
	c := make(chan int, 3)
	go produzir(c)

	//Causamos um deadlock, pois o quarto valor nunca foi produzido.
	fmt.Println(<-c, <-c, <-c, <-c)
}

func produzir(c chan int) {
	c <- 1
	c <- 2
	c <- 3

	//Com o "close" o canal é fechado e indica que não será mais enviado nenhum valor
	//close(c)
}