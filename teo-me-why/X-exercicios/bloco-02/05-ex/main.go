package main

import "fmt"

func main() {

	fmt.Println("Informe um número para ser dobrado:")

	var numero float64
	fmt.Scanf("%f", &numero)

	fmt.Printf("O dobro de%2.f é%2.f\n", numero, numero*2)
}
