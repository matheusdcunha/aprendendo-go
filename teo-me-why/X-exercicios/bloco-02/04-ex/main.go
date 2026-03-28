// Faça um programa que receba um número inteiro e calcule sua raiz quadrada e exiba o resultado.

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Digite um número inteiro que deseja saber a raiz quadrada:")

	var numero int
	fmt.Scanf("%f", &numero)

	raizQuadrada := math.Sqrt(float64(numero))

	fmt.Printf("A raiz quadrada de %d é %.2f\n", numero, raizQuadrada)
}
