//Escreva um programa que solicite ao usuário um número e exiba a tabuada desse número de 1 a 10.

package main

import (
	"fmt"
	"strconv"
)

func tabuada(n int) {
	fmt.Printf("| Tabuada de %02d |\n", n)
	for i := range 10 {
		multi := (i + 1) * n
		fmt.Printf("| %02d x %02d = %02d |\n", i, n, multi)
	}
}

func main() {
	var input string

	fmt.Printf("Informe o número que deseja saber a tabuada: ")
	fmt.Scanf("%s", &input)
	fmt.Println("")

	n, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("Infome um número inteiro válido")

	} else {
		tabuada(n)
	}
}
