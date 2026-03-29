// Faça o programa de uma sorveteria, onde o usuário pode escolher:
// Tipo de sorvete: casquinha (R$1,00), cascão (R$2,50), cestinha (R$4,00)
// Sabor do sorvete: morango, creme, chocolate
// Cobertura: Caramelo (R$1,50), morango (R$1,50), chocolate (R$1,50), sem cobertura (R$0,00).
// Apresente o valor a ser pago

package main

import "fmt"

func main() {
	var valor float64
	var tipoDeSorvete string

	fmt.Println(`
Escolha o tipo de sorvete
---
1 - Casquinha(R$ 1,00)
2 - Cascão(R$ 2,50)
3 - Cestinha(R$ 4,00)
---`)

	var tipoDeSorveteOpcao int
	fmt.Scanf("%d", &tipoDeSorveteOpcao)

	switch tipoDeSorveteOpcao {
	case 1:
		tipoDeSorvete = "Casquinha"
		valor += 1.00
	case 2:
		tipoDeSorvete = "Cascão"
		valor += 2.5
	case 3:
		tipoDeSorvete = "Cestinha"
		valor += 4.0
	}

	var sabor string

	fmt.Println(`
Escolha o sabor do sorvete
---
1 - Morango
2 - Creme
3 - Chocolate
---`)

	var saborOpcao int
	fmt.Scanf("%d", &saborOpcao)

	switch saborOpcao {
	case 1:
		sabor = "Morango"
	case 2:
		sabor = "Creme"
	case 3:
		sabor = "Chocolate"
	}

	var cobertura string

	fmt.Println(`
Escolha a cobertura
---
1 - Caramelo(R$ 1,50)
2 - Morango(R$ 1,50)
3 - Chocolate(R$ 1,50)
4 - Sem cobertura(R$ 0,00)
---`)

	var coberturaOpcao int
	fmt.Scanf("%d", &coberturaOpcao)

	switch coberturaOpcao {
	case 1:
		cobertura = "Caramelo"
		valor += 1.50
	case 2:
		cobertura = "Morango"
		valor += 1.50
	case 3:
		cobertura = "Chocolate"
		valor += 1.50
	case 4:
		cobertura = "Sem cobertura"
	}

	if cobertura == "Sem cobertura" || cobertura == "" {
		fmt.Printf("Você escolheu uma %s de %s e sem cobertura\n", tipoDeSorvete, sabor)
		fmt.Println("O total a ser pago foi:", valor)
	} else {
		fmt.Printf("Você escolheu uma %s de %s e a cobertura de %s\n", tipoDeSorvete, sabor, cobertura)
		fmt.Println("O total a ser pago foi: R$", valor)
	}

}
