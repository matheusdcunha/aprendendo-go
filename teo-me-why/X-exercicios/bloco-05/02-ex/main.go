// Faça um programa que receba 6 temperaturas. Remova a 1a e a última para calcular a média. Se a média for acima de 30 graus, exiba que houve um aumento da temperatura.

package main

import (
	"fmt"
	"strconv"
)

func main() {
	temperaturas := [6]float64{}

	for i := range temperaturas {
		fmt.Printf("Informe a %dª temperatura\n", i+1)
		var input string
		fmt.Scanf("%s", &input)

		temperatura, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Informe uma temperatura válida!")
		}

		temperaturas[i] = temperatura
	}

	temperaturasParaMedia := temperaturas[1 : len(temperaturas)-1]

	var media float64
	for _, valor := range temperaturasParaMedia {
		media += valor
	}

	media = media / float64(len(temperaturasParaMedia))

	if media > 30 {
		fmt.Println("Houve um aumento da temperatura")
	} else {
		fmt.Println("Não houve um aumento da temperatura")
	}
}
