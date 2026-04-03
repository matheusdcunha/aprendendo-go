// Faça um programa que receba dois valores A e B. Faça a potência desses dois valores e retorne o resultado:

// a ^ b = z

package main

import (
	"fmt"
	"math"
	"strconv"
)

func potencia(a, b float64) float64 {
	return math.Pow(a, b)
}

func main() {
	valores := [2]float64{}

	for i := range valores {

		fmt.Printf("Digite o %d valor: ", i+1)
		var input string
		fmt.Scanf("%s", &input)

		a, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Informe um valor válido!")
			break
		} else {
			valores[i] = a
		}
	}

	if valores[1] != 0 {
		fmt.Printf("A potência dos valores é: %.2f", potencia(valores[0], valores[1]))
	}
}
