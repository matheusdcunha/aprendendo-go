package main

import "fmt"

type Smarthphone struct {
	Marca         string
	Armazenamento uint
	Cor           string
	Peso          float64
	Valor         float64
	Ligado        bool
}

func (s *Smarthphone) Ligar() {
	s.Ligado = true
}

func (s *Smarthphone) Desligar() {
	s.Ligado = false
}

type Pessoa struct {
	Nome  string
	Idade uint
}

type Funcionario struct {
	Pessoa
	Cargo string
}

func (f Funcionario) Trabalho() {
	fmt.Printf("O %s trabalha de %s\n", f.Nome, f.Cargo)
}

type Exemplo struct {
	Valores [3]int
}

func main() {

	a := Exemplo{Valores: [3]int{1, 2, 3}}
	b := Exemplo{Valores: [3]int{1, 2, 3}}

	fmt.Println("A = B", a == b)

	celular := Smarthphone{Marca: "Apple", Armazenamento: 64, Cor: "Branco", Peso: 150.30, Valor: 2000.12}

	celular.Ligar()

	if celular.Ligado {
		fmt.Printf("O %s está ligado\n", celular.Marca)
	} else {
		fmt.Printf("O %s está desligado\n", celular.Marca)
	}

	funcionario := Funcionario{
		Pessoa: Pessoa{Nome: "Matheus", Idade: 27},
		Cargo:  "Dev",
	}

	funcionario.Trabalho()
}
