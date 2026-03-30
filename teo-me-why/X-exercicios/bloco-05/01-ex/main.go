// Faça um programa que receba 4 notas, calcule a média, mínimo e máximo dessas notas.

package main

import (
	"fmt"
	"slices"
	"strconv"
)

func main() {

	notas := [4]float64{}

	for i := range notas {
		fmt.Printf("Informe a %dª nota\n", i+1)
		var input string
		fmt.Scanf("%s", &input)

		nota, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Informe uma nota válida!")
		}

		notas[i] = nota
	}

	var media float64

	minima := slices.Min(notas[:])
	maxima := slices.Max(notas[:])

	for _, valor := range notas {
		media += valor
	}

	media = media / float64(len(notas))

	fmt.Printf(`---
Média: %f
Mínima: %f
Máxima: %f
---`, media, minima, maxima)

}
