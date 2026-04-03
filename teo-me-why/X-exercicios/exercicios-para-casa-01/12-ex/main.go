//Faça um programa que receba um número e exiba seu fatorial. Use função para isso.

package main

import (
	"fmt"
	"strconv"
)

func fatorial(n int) (res int) {
	res = 1

	if n == 0 || n == 1 {
		return
	}

	for i := range n {
		res *= i + 1
	}

	return
}

func main() {
	var input string

	fmt.Printf("Informe o número que deseja saber o fatorial: ")
	fmt.Scanf("%s", &input)

	n, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("Infome um número inteiro válido")

	} else {
		fmt.Println("---")
		fmt.Printf("O fatorial de %02d é: %02d\n", n, fatorial(n))
		fmt.Println("---")
	}
}
