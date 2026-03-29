// Faça um programa que receba uma quantidade indefinida de valores correspondentes a “saldo em conta”, mas quando o usuário apertar “enter” sem digitar valor algum, o programa para de receber valores, e exibe a soma de todos os valores digitados anteriormente.

package main

import (
	"fmt"
	"strconv"
)

func main() {

	var valorTotal float64

	for i := 0; true; i++ {
		fmt.Printf("Digite o %dª valor:\n", i+1)

		var input string
		fmt.Scanf("%s", &input)

		if input == "" {
			break
		}

		valor, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Entre com um valor válido!!")
			i--
			continue
		}

		valorTotal += valor
	}

	fmt.Println("O valor total é: R$", valorTotal)
}
