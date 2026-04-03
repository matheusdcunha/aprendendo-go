//Escreva um programa com a função “quadrado”, que recebe um número inteiro e eleva ele ao quadrado, retornando o resultado.

package main

import "fmt"

func quadrado(n float64) (res float64) {
	res = n * n
	return
}

func main() {
	n := float64(4)
	nQuadrado := quadrado(4)

	fmt.Println("Número:", n)
	fmt.Println("Número ao quadrado:", nQuadrado)
}
