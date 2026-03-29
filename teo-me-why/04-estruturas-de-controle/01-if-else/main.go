package main

import "fmt"

func main() {

	var idade int

	fmt.Println("Qual sua idade?")
	fmt.Scanf("%d", &idade)

	if idade >= 18 {
		fmt.Println("Você é um adulto")
	} else if 14 <= idade {
		fmt.Println("Você é um adolescente")
	} else {
		fmt.Println("Você é uma criança")
	}
}
