// Faça um programa que use uma função para responder a saudação de um usuário: “Olá, fulano! Boas vindas!”. Use funções para isso.

package main

import "fmt"

func boasVindas(nome string) {
	fmt.Printf("Olá, %s! Boas vindas!\n", nome)
}

func main() {
	fmt.Println("Digite seu nome:")
	var nome string
	fmt.Scanf("%s", &nome)

	boasVindas(nome)
}
