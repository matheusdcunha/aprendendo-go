// Faça um programa que receba um número em segundos, converta esse número para horas, minuto e segundos. Exemplos:

// Entrada: 556
// Saída: 0:9:16

// Entrada: 140153
// Saída: 38:55:53

package main

import (
	"fmt"
	"strconv"
	"time"
)

func convertToHours(n int) (res string) {
	duration := time.Duration(n) * time.Second

	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60

	res = fmt.Sprintf("%02dh:%02dm:%02ds\n", hours, minutes, seconds)

	return
}

func main() {
	fmt.Printf("Digite a quantidade de segundos: ")
	var input string
	fmt.Scanf("%s", &input)

	seconds, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("Digite um valor válido!")
	} else {
		fmt.Println(convertToHours(seconds))
	}
}
