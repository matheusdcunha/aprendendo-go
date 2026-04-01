package main

import (
	"fmt"
	"os"
)

func sum(a, b int) int {
	return a + b
}

func teste(a, b int, c string) (valor int, nome string) {

	valor = a + b
	nome = c
	return
}

func media(a, b int) (float64, error) {
	total := sum(a, b)
	res := float64(total) / float64(2)
	err := fmt.Errorf("Deu erro")

	return res, err
}

func main() {

	valor, nome := teste(1, 2, "Laís")
	fmt.Println(valor)
	fmt.Println(nome)

	med, err := media(4, 6)
	if err != nil {
		fmt.Println("Error!", err)
		os.Exit(1)
	}
	fmt.Println("Média", med)
}
