//Escreva um programa que solicite ao usuário frases. Para parar de solicitar frases, ele pode apenas apertar o “enter”.
// Seu programa deve apresentar cada frase e quantas vezes ela foi repetida.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	frases := map[string]int{}

	for {
		fmt.Printf("Informe uma frase: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		frase := strings.TrimSpace(scanner.Text())

		if frase == "" {
			break
		}

		frases[frase]++
	}

	fmt.Println()

	for k, v := range frases {
		fmt.Printf("A frase %q aparece %02d vezes\n", k, v)
	}

}
