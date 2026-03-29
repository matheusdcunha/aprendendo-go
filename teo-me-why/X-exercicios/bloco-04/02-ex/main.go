// Faça um programa que receba 4 alturas usando um laço de repetição e realize a soma dessas alturas.

package main

import "fmt"

func main() {

	var altura, somaAlturas float64

	for i := range 4 {
		fmt.Printf("Digite a %dª altura:\n", i+1)

		fmt.Scanf("%f", &altura)
		somaAlturas += altura
	}

	fmt.Println("A soma das alturas é:", somaAlturas)
}
