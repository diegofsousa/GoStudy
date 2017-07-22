package main

import (
	"fmt"
)

type Estado struct{
	nome string
	populacao int
	capital string
}

func main() {
	estados := make(map[string]Estado, 6)
	estados["GO"] = Estado{"Goiás", 6434052, "Goiânia"}
	estados["PB"] = Estado{"Paraíba", 3914418, "João Pessoa"}
	estados["PR"] = Estado{"Paraná", 10997462, "Curitiba"}
	estados["AM"] = Estado{"Amazonas", 3807923, "Manaus"}
	estados["RN"] = Estado{"Rio Grande do Norte", 3373960, "Natal"}
	estados["SE"] = Estado{"Sergipe", 2228489, "Aracaju"}

	fmt.Println(estados)

	/*Recuperando itens na lista*/

	sergipe := estados["SE"]
	fmt.Println(sergipe)

	fmt.Println(estados["SP"])

	saoPaulo, encontrado := estados["SP"]
	if encontrado {
		fmt.Println(saoPaulo)
	}

	/*Deletando itens*/
	delete(estados, "AM")
	fmt.Println(estados)

	/*Iterando um mapa*/

	for sigla, estado := range estados {
		fmt.Printf("%s (%s) possui %d habitantes.\n",
			estado.nome, sigla, estado.populacao)
	}

}

// Paramos na ṕágina 63 - Cap. 5 - Criando novos tipos