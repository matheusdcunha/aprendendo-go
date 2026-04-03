// Faça um programa que receba um número. Este número corresponde a uma posição na sequência de Fibonacci: 1, 1, 2, 3, 5,...

// Exiba o número da sequência cuja posição foi informada:
// 	A posição x corresponde ao número y

package main

import (
	"fmt"
	"strconv"
)

func fibonacci(n int) int {
	if n <= 2 {
		return 1
	}

	a, b := 1, 1
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}

	return b
}

func main() {
	var input string

	fmt.Printf("Informe o número que deseja saber sua posição em fibonnaci: ")
	fmt.Scanf("%s", &input)

	n, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("Infome um número inteiro válido")

	} else {
		fmt.Printf("A posição %d corresponde ao número %d\n", n, fibonacci(n))
	}
}
