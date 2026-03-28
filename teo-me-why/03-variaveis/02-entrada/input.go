package main

import "fmt"

func main() {

	fmt.Println("Digita seu nome ai:")

	var nome string
	fmt.Scanf("%s", &nome)

	fmt.Println("Seu nome é:", nome)

	fmt.Println("Digite um valor monetário:")

	var valorMonetario int
	fmt.Scanf("%d", &valorMonetario)

	fmt.Println("O valor é:", valorMonetario)
}
