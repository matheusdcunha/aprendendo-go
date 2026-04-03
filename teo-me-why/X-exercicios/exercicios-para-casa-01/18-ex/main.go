// Escreva um programa que solicite ao usuário uma palavra e verifique se a palavra é um palíndromo (ou seja, é a mesma palavra quando lida de trás para frente).

package main

import (
	"fmt"
	"strings"
)

func palindromo(p string) (res bool) {

	p = strings.ToLower(p)
	runes := []rune(p)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}

func main() {
	var input string

	fmt.Printf("Informe uma palavra para sabe se é palindromo: ")
	fmt.Scanf("%s", &input)

	if palindromo(input) {
		fmt.Printf("%q é palindromo\n", input)
	} else {
		fmt.Printf("%q não é palindromo\n", input)
	}

}
