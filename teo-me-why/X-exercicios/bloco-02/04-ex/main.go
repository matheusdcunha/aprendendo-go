// Faça um programa que receba um número inteiro e calcule sua raiz quadrada e exiba o resultado.

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Digite um número inteiro que deseja saber a raiz quadrada:")

	var numero float64
	fmt.Scanf("%f", &numero)

	raizQuadrada := math.Sqrt(numero)

	fmt.Printf("A raiz quadrada de %f é %f\n", numero, raizQuadrada)
}
