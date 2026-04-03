// Escreva um programa que receba uma lista de números do usuário e conte quantas vezes um número específico aparece na lista. Solicite ao usuário um número e exiba a contagem.

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Informe uma lista de números inteiros")

	contagem := map[int]int{}

	i := 1
	for {
		var input string

		fmt.Printf("%02dº número: ", i)
		fmt.Scanf("%s", &input)

		if input == "" {
			break
		}

		n, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Infome um número inteiro válido")
			continue
		}

		contagem[n]++
		i++
	}

	var input string

	fmt.Printf("Digite um número inteiro para contagem: ")
	fmt.Scanf("%s", &input)

	numero, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Infome um número inteiro válido")
	} else {
		fmt.Printf(`---
O número %d aparece %d vezes
---
`, numero, contagem[numero])
	}

}
