package main

import "fmt"

func main() {
	idades := make(map[string]int)

	idades["Matheus"] = 27
	idades["Laís"] = 29

	teste := map[int]bool{
		1: false,
		2: true,
	}

	valor, prs := idades["Guarana"]
	if prs {
		fmt.Println("Idade Guarana:", valor)
	} else {
		fmt.Println("Não encontrei")
	}

	cursos := map[string][]string{
		"Matheus": {"Kubernetes", "Go"},
		"Laís":    {"Cinema", "Cerâmica"},
	}

	fmt.Println(idades)
	fmt.Println(teste)
	fmt.Println(cursos)
}
