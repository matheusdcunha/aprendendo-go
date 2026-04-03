// Faça um programa que receba o raio de uma circunferência em centímetros. Retorne para o usuário qual é a área e perímetro desta circunferência no seguinte formato.

// Área:  x.xx
// Perímetro:  y.yy

package main

import (
	"fmt"
	"math"
	"strconv"
)

func calcArea(raio float64) float64 {
	raioQuadrado := raio * raio

	return math.Pi * raioQuadrado
}

func calcPerimetro(raio float64) float64 {
	piDobrado := math.Pi * 2

	return piDobrado * raio
}

func main() {
	fmt.Printf("Digite o raio em centímetros: ")
	var input string
	fmt.Scanf("%s", &input)

	raio, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Informe um valor válido!")
	}

	fmt.Printf("A área da circuferência é : %.2f cm\n", calcArea(raio))
	fmt.Printf("O perímetro da circuferência é : %.2f cm\n", calcPerimetro(raio))

}
