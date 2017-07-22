package main

import (
	"errors"
	"fmt"
)

func main() {
	pilha := Pilha{} // Instância do objeto pilha
	fmt.Println("Pilha criada com tamanho ", pilha.Tamanho()) // Todos os métodos em Go devem estar com a inicial maiúscula
	fmt.Println("Vazia? ", pilha.Vazia())

	pilha.Empilhar("Go")
	pilha.Empilhar(2009)
	pilha.Empilhar(3.14)
	pilha.Empilhar(true)

	fmt.Println("Tamanho após empilhamento: ", pilha.Tamanho())
	fmt.Println("Vazia? ", pilha.Vazia())

	// Em Go só existe o "for" para estrutura de repetição. Em compensação é um puta "for"! Este por exemplo, simula um "while".
	for !pilha.Vazia(){
		v, _ := pilha.Desempilhar()
		fmt.Println("Desempilhando ", v)
		fmt.Println("Tamanho: ", pilha.Tamanho())
		fmt.Println("Vazia? ", pilha.Vazia())
	}

	_, err := pilha.Desempilhar()
	if err != nil {
		fmt.Println("Erro: ", err)
	}
}

// Não existem "class" em Go. Structs fazem este papel
type Pilha struct {
	valores []interface{} // O tipo interface declara que o mesmo pode ser qualquer tipo: int, string, bool.
}

/*
Para saber que esta função é relativa a Pilha
devemos declarar no início com o nome da struct relacionada.
Em Go não existem receptorem implicitos: this ou self. */
func (pilha Pilha) Tamanho() int{
	return len(pilha.valores)
}

func (pilha Pilha) Vazia() bool{
	return pilha.Tamanho() == 0
}

/*
No caso dos métodos Empilhar() e Desempilhar() , desejamos al-
terar a pilha na qual tais métodos foram chamados. Em Go, argumentos de
funções e métodos são sempre passados por cópia (com a exceção de slices,
maps e channels, que são conhecidos como reference types, como veremos
nos próximos capítulos). Por isso, quando precisamos alterar qualquer argu-
mento – incluindo receptores de métodos – devemos declará-los como pon-
teiros. Vejamos a definição do método Empilhar()

ref: Programando em Go - Casa do Código.
*/


func (pilha *Pilha) Empilhar(valor interface{}) {
	pilha.valores = append(pilha.valores, valor)
}

func (pilha *Pilha) Desempilhar() (interface{}, error) {
	if pilha.Vazia() {
		return nil, errors.New("Pilha vazia!")
	}
	valor := pilha.valores[pilha.Tamanho()-1]
	pilha.valores = pilha.valores[:pilha.Tamanho()-1]
	return valor, nil
}

/*Paramos na página 39 - Arrays*/