// Faça um programa que vende uma garrafa de água:
// Se o cliente escolher água mineral natural, será cobrado R$1,50
// Se o cliente escolher água mineral com gás, será cobrado R$2,50

package main

import "fmt"

func main() {

	fmt.Println(`
Escolha o tipo de água
---
1 - Água Mineral Natural(R$ 1,50)
2 - Água Mineral com Gás(R$ 2,50)
---`)

	var escolha int
	fmt.Scanf("%d", &escolha)

	switch escolha {
	case 1:
		fmt.Println("Valor a ser pago: R$ 1,50")
	case 2:
		fmt.Println("Valor a ser pago: R$ 2,50")
	default:
		fmt.Println("Opção inválida")
	}
}
