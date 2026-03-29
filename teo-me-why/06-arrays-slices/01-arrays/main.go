package main

import "fmt"

func main() {

	var a = [2]float64{1.12, 2.76}
	numeros := [5]int{1, 2, 3, 4, 5}
	nomes := [5]string{"Laís", "Matheus", 4: "Guaraná"}

	matriz := [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}

	var valorTotalNumeros int

	for _, valor := range numeros {
		valorTotalNumeros += valor
	}

	fmt.Println("Valor do array a:", a)
	fmt.Println("Array de nomes:", nomes)
	fmt.Println("Array multidimensional:", matriz)
	fmt.Println("Valor total do arrays de números é:", valorTotalNumeros)
	fmt.Println("Dois primerios números do array de números:", numeros[0:2])
	fmt.Println("Dois últimos números do array de números:", numeros[len(numeros)-2:])
}
