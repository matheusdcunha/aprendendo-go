package main

import "fmt"

func main() {
	fmt.Println(`
Opções de escolha
---
1 - Banan
2 - Maça
3 - Manga
4 - Uva`)

	var opcao int
	fmt.Scanf("%d", &opcao)

	switch opcao {
	case 1:
		fmt.Println("Banana")
	case 2:
		fmt.Println("Maça")
	case 3:
		fmt.Println("Manga")
	case 4:
		fmt.Println("Uva")
	default:
		fmt.Println("Opção não existente")
	}

}
