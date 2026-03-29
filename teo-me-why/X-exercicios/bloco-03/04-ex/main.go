// Faça um programa que verifique se a pessoa pertence à família “calvo”.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(`---
Digite seu nome completo
---`)

	nome, _ := reader.ReadString('\n')

	nome = strings.TrimSpace(nome)

	if strings.Contains(strings.ToLower(nome), "calvo") {
		fmt.Printf("%s você é da familia Calvo", nome)
	} else {
		fmt.Printf("%s você NÃO é da familia Calvo", nome)
	}

}
