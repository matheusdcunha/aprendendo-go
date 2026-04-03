// Escreva um programa que solicite ao usuário seu nome e depois seu sobrenome.

// Exiba ambos depois utilizando apenas uma variável.

package main

import "fmt"

func main() {

	usuario := map[string]string{
		"nome":      "",
		"sobrenome": "",
	}

	var input string

	fmt.Printf("Digite seu nome: ")
	fmt.Scanf("%s", &input)

	usuario["nome"] = input

	fmt.Printf("Digite seu sobrenome: ")
	fmt.Scanf("%s", &input)

	usuario["sobrenome"] = input

	fmt.Printf("Seu nome completo é: %s %s", usuario["nome"], usuario["sobrenome"])
}
