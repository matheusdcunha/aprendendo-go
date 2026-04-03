// Faça um programa que calcule a média de uma quantidade indefinida de valores, usando uma função para isso.

package main

import (
	"fmt"
	"strconv"
)

func calcMedia(values ...float64) float64 {
	var total float64

	for _, v := range values {
		total += v
	}
	return total / float64(len(values))
}

func getValores() (values []float64) {

	for {
		var valorStr string
		fmt.Printf("Entre com o valor: ")
		fmt.Scanf("%s", &valorStr)

		if valorStr == "" {
			return values
		}

		valor, err := strconv.ParseFloat(valorStr, 64)

		if err != nil {
			fmt.Println("Entre com um valor válido!")
			continue
		}

		values = append(values, valor)
	}
}

func main() {

	valores := getValores()
	media := calcMedia(valores...)

	fmt.Printf("Média: %.2f", media)
}
