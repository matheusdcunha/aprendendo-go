// Faça um programa que verifique se o item que a pessoa escolheu para comprar na loja está na lista: laranja, cerveja, miojo, carvão, picanha.

package main

import (
	"fmt"
	"strings"
)

func main() {

	var opcao string
	fmt.Println("Entre com o item:")

	fmt.Scanf("%s", &opcao)

	itens := "laranja, cerveja, miojo, carvão, picanha"

	if strings.Contains(itens, strings.ToLower(opcao)) {
		fmt.Println("Produto válido para compra")
	} else {
		fmt.Println("Produto não disponível para compra")
	}
}
