// Faça um programa que receba um número. Verifique se o número informado é par ou ímpar. Exiba o resultado da seguinte maneira:

// 	O número x é impar
// ou
// 	O número x é par

package main

import (
	"fmt"
	"os"
	"strconv"
)

func ehPar(a int) bool {
	return a%2 == 0
}

func main() {
	fmt.Printf("Digite um número: ")
	var input string
	fmt.Scanf("%s", &input)

	numero, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("Digite um valor válido!")
		os.Exit(1)
	}

	par := ehPar(numero)

	if par {
		fmt.Printf("O número %d é par\n", numero)
	} else {
		fmt.Printf("O número %d é impar\n", numero)
	}
}
