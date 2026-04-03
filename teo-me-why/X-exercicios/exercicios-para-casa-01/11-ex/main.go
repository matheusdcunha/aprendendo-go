// Escreva um programa que receba uma lista de números do usuário e conte quantas vezes um número específico aparece na lista. Solicite ao usuário um número e exiba a contagem.

package main

import (
	"fmt"
	"strconv"
)

func contagem(lista []int, numero int) (res int) {
	res = 0

	for _, v := range lista {
		if v == numero {
			res++
		}
	}

	return
}

func main() {
	fmt.Println("Informe uma lista de números inteiros")

	lista := []int{}

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

		lista = append(lista, n)
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
`, numero, contagem(lista, numero))
	}

}
