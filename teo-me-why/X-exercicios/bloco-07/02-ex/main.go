// Faça um programa que informe se uma número é par ou ímpar com uma função que se chama “EhPar”, retornando um booleano.

package main

import "fmt"

func ehPar(n int) bool {
	return n%2 == 0
}

func main() {
	fmt.Println(ehPar(2003))
}
