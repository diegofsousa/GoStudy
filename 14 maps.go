package main

import (
	"fmt"
)

func main() {
	// declaração de maps
	//vazio1 := map[int]string{}
	//vazio2 := make(map[int]string)
	//mapaGrande := make(map[int]string, 4096)

	capitais := map[string]string{
		"GO":"Goiania",
		"PB":"João Pessoa",
		"PR":"Curitiba"}

	fmt.Println(len(capitais))

	capitais["RN"] = "Natal"
	capitais["AM"] = "Manaus"
	capitais["SE"] = "Aracaju"

	fmt.Println(capitais)

	populacao := make(map[string]int, 3)
	populacao["GO"] = 6434052
	populacao["PB"] = 3914418
	populacao["PR"] = 10997462

	fmt.Println(populacao)

}