package main

import (
	"fmt"
	"strconv"
)

func validaInput() float64 {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recuperado: %v", r)
		}
	}()

	var input string
	fmt.Println("Entre com um float:")
	fmt.Scanf("%s", &input)

	num, err := strconv.ParseFloat(input, 64)
	if err != nil {
		panic("entrada inválida")
	}

	return num
}

func main() {

	numero := validaInput()
	fmt.Println(numero)
}
