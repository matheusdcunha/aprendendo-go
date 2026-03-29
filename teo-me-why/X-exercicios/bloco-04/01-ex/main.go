// Faça um programa que conte quantas vezes a letra “a” aparece em uma palavra

package main

import (
	"fmt"
	"strings"
)

func main() {
	var contador int

	fmt.Println("PROGRAMA PARA CALCULAR QUANTIDADE DE A's")
	fmt.Println(`---
Digite uma palavra:
---`)

	var palavra string
	fmt.Scanf("%s", &palavra)

	palavraLower := strings.ToLower(palavra)

	for _, valor := range palavraLower {

		if string(valor) == "a" {
			contador++
		}

	}

	fmt.Printf(`---
A palavra %s contém %d letra A`, palavra, contador)
}
