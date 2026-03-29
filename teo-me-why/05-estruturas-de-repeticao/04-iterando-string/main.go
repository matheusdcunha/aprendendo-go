package main

import "fmt"

func main() {
	nome := "Matheus"

	for i := range nome {
		letra := string(nome[i])
		fmt.Println(i+1, letra)
	}

	fmt.Println("---Com Bytes---")
	for i, valor := range nome {
		letra := string(valor)
		fmt.Println(i+1, letra)
	}
}
