// Solicite ao usuário o nome de uma fruta e exiba o preço correspondente e some ao total de uma compra.
// Quando o usuário inserir uma string vazia, encerre a compra e informe o valor total da compra.

package main

import (
	"fmt"
	"strings"
)

func main() {

	frutas := map[string]float64{
		"maçã":    1.50,
		"banana":  2.75,
		"uva":     1.90,
		"pera":    1.25,
		"laranja": 0.65,
		"limão":   1.2,
		"goiaba":  2.15,
		"abacaxi": 3.20,
		"jaca":    5.80,
	}

	total := 0.0

	for {
		var fruta string
		fmt.Printf("Digite a fruta que quer comprar: ")
		fmt.Scanf("%s", &fruta)

		if fruta == "" {
			break
		}

		valor, ok := frutas[strings.ToLower(fruta)]
		if !ok {
			fmt.Printf(`----
%s não disponível para compra
----
`, fruta)
			continue
		}
		fmt.Printf(`----
%s adicionada ao carrinho
----
`, fruta)
		total += valor

	}

	fmt.Printf("O valor total da sua compra foi: %.2f reais", total)

}
