// Escreva um programa que receba o nome e a idade de uma pessoa. Depois exiba a mensagem:

// “Olá fulano, bom saber que você tem x anos. Boas vindas!”

package main

import "fmt"

func boasVindas(nome string, idade uint8) {
	fmt.Printf("Olá %s, bom saber que você tem %d anos. Boas vindas!", nome, idade)
}

func main() {
	fmt.Printf("Digite seu nome: ")
	var nome string
	fmt.Scanf("%s", &nome)

	fmt.Printf("Digite sua idade: ")
	var idade uint8
	fmt.Scanf("%d", &idade)

	boasVindas(nome, idade)
}
