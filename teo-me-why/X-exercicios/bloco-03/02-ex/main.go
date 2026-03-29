// Faça um programa que vende uma garrafa de água:
// Se o cliente escolher água mineral natural, será cobrado R$1,50
// Se o cliente escolher água mineral com gás, será cobrado R$2,50
// considerar a quantidade de garrafas de água

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

	var quantidade int
	fmt.Println(`---
Digite a quantidade desejada:
---`)
	fmt.Scanf("%d", &quantidade)

	switch escolha {
	case 1:
		fmt.Println("Valor a ser pago:", 1.50*float64(quantidade))
	case 2:
		fmt.Println("Valor a ser pago:", 2.40*float64(quantidade))
	default:
		fmt.Println("Opção inválida")
	}
}
