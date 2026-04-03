//Considere a seguinte lista:
// [123, 435, 987, 1984, 2, 19, 423, -178, 320]

// Faça um programa que retorne a posição do menor e do maior valor encontrado:

// O maior valor está na posição x
// O menor valor está na posição y

package main

import "fmt"

func main() {

	lista := [9]int{123, 435, 987, 1984, 2, 19, 423, -178, 320}

	maior := map[string]int{"valor": 0, "pos": 0}
	menor := map[string]int{"valor": 0, "pos": 0}

	for pos, valor := range lista {
		if valor > maior["valor"] {
			maior["valor"] = valor
			maior["pos"] = pos
		}

		if valor < menor["valor"] {
			menor["valor"] = valor
			menor["pos"] = pos
		}
	}

	fmt.Printf("O maior valor está na posição %d\n", maior["pos"])
	fmt.Printf("O menor valor está na posição %d\n", menor["pos"])

}
