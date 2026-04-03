//Faça um programa que receba um número. Verifique se este número é primo ou não, e retorne o resultado:

// O número x é primo
// ou
// 	O número x não é primo

package main

import (
	"fmt"
	"strconv"
)

func ehPrimo(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var input string

	fmt.Printf("Informe o número que deseja saber se é primo: ")
	fmt.Scanf("%s", &input)

	n, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("Infome um número inteiro válido")

	} else {
		if ehPrimo(n) {
			fmt.Printf("O número %02d é primo", n)
		} else {
			fmt.Printf("O número %02d é não é primo", n)
		}
	}
}
