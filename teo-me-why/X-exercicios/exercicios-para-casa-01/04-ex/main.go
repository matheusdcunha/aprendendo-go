// Faça um programa que receba dois valores A e B. Faça a soma desses dois valores e retorne o resultado:

// Soma:  x.xx

package main

import (
	"fmt"
	"strconv"
)

func sum(a, b float64) float64 {
	return a + b
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
		fmt.Printf("A soma dos valoers é: %.2f", sum(valores[0], valores[1]))
	}

}
