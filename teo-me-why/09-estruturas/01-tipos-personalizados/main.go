package main

import (
	"fmt"
	"strconv"
)

type Altura float64
type Peso float64
type Nome string

func main() {

	peso := Peso(70)
	altura := Altura(1.70)
	nome := Nome("12.1")

	fmt.Printf("Peso %v\n", peso)
	fmt.Printf("Altura %v\n", altura)

	imc := float64(peso) / float64(altura*altura)

	fmt.Printf("IMC: %.2f\n", imc)

	nomeFloat, err := strconv.ParseFloat(string(nome), 64)
	if err != nil {
		fmt.Println("Deu errado")
	}

	fmt.Println("Nome em float", nomeFloat)
}
