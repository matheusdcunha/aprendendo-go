//Escreva um programa que receba o nome de uma pessoa e faça uma saudação.

// “Olá fulano! Boas vindas!”

package main

import "fmt"

func boasVindas(nome string) {
	fmt.Printf("Olá %s! Boas vindas!", nome)
}

func main() {
	fmt.Printf("Digite seu nome: ")
	var nome string
	fmt.Scanf("%s", &nome)

	boasVindas(nome)
}
