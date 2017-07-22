package main

import (
	"fmt"
)

/*Neste caso o retorno é declarado no início da função*/
func PrecoFinal(precoCusto float64) (precoDolar float64, precoReal float64){
	fatorLucro := 1.33
	taxaConversao := 2.34
	//usa-se "=", pois estas variávei já foras declaradas.
	precoDolar = precoCusto * fatorLucro
	precoReal = precoDolar * taxaConversao

	return
}

func main() {
	precoDolar, precoReal := PrecoFinal(34.99)
	fmt.Printf("Preço final em dólar: %.2f\n"+
		"Preço final em reais: %.2f\n",
		precoDolar, precoReal)
}