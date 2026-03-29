// Faça um programa que a pessoa entra com 3 notas e verificamos se a média delas é suficiente para passar na prova. Nota de corte 6

package main

import "fmt"

func main() {

	var notaTotal float64
	var notaAuxiliar float64

	fmt.Println("Digite a primeira nota:")

	fmt.Scanf("%f", &notaAuxiliar)
	notaTotal += notaAuxiliar

	fmt.Println("Digite a segunda nota:")

	fmt.Scanf("%f", &notaAuxiliar)
	notaTotal += notaAuxiliar

	fmt.Println("Digite a terceira nota:")

	fmt.Scanf("%f", &notaAuxiliar)
	notaTotal += notaAuxiliar

	fmt.Println("A média é:", notaTotal/3)
}
