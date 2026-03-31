// Faça o programa de uma sorveteria, onde o usuário pode escolher:
// Tipo de sorvete: casquinha (R$1,00), cascão (R$2,50), cestinha (R$4,00)
// Sabor do sorvete: morango, creme, chocolate
// Cobertura: Caramelo (R$1,50), morango (R$1,50), chocolate (R$1,50), sem cobertura (R$0,00).
// Apresente o valor a ser pago.

package main

import (
	"fmt"
	"os"
)

func main() {

	tipoSorvete := map[string]float64{
		"casquinha": 1.0,
		"cascão":    2.5,
		"cestinha":  4.0,
	}

	saborSorvete := map[string]float64{
		"morango":   0,
		"creme":     0,
		"chocolate": 0,
	}

	coberturaSorvete := map[string]float64{
		"caramelo":  1.5,
		"morango":   1.5,
		"chocolate": 1.5,
	}

	itens := map[string]map[string]float64{
		"tipos":      tipoSorvete,
		"sabores":    saborSorvete,
		"coberturas": coberturaSorvete,
	}

	var tipo, sabor, cobertura string

	fmt.Printf("Escolha um tipo [casquinha/cascão/cestinha]: ")
	fmt.Scanf("%s", &tipo)

	fmt.Printf("Escolha um tipo [morango/creme/chocolate]: ")
	fmt.Scanf("%s", &sabor)

	fmt.Printf("Escolha um tipo [caramelo/morango/chocolate]: ")
	fmt.Scanf("%s", &cobertura)

	total := 0.0

	if valor, ok := itens["tipos"][tipo]; !ok {
		fmt.Println("Tipo inválido")
		os.Exit(1)
	} else {
		total += valor
	}

	if valor, ok := itens["sabores"][tipo]; !ok {
		fmt.Println("Sabor inválido")
		os.Exit(1)
	} else {
		total += valor
	}

	if valor, ok := itens["coberturas"][tipo]; !ok {
		fmt.Println("Cobertura inválida")
		os.Exit(1)
	} else {
		total += valor
	}

	fmt.Println("Total: R$", total)

}
