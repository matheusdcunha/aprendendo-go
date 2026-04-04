package main

import "fmt"

type Personagem struct {
	Nome   string
	Raca   string
	Classe string
	Nivel  int
	Atacante
}

type Jogador struct {
	Personagem
	Username string
}

func (j *Jogador) SubirNivel() {
	j.Nivel++
}

func (j Jogador) ForcaAtaque() int {
	return j.Nivel * 10
}

type Inimigo struct {
	Personagem
	Tipo string
}

func (i Inimigo) ForcaAtaque() int {
	return i.Nivel + 10
}

type NPC struct {
	Personagem
	Funcao string
}

type Atacante interface {
	ForcaAtaque() int
}

func calcForcaAtaque(a Atacante) int {
	return a.ForcaAtaque()
}

func showTipo(a Atacante) {
	switch v := a.(type) {
	case Jogador:
		fmt.Println("Jogador:", v.Username)
	case Inimigo:
		fmt.Println("Inimigo:", v.Nome)
	}
}

func main() {
	j := Inimigo{Personagem: Personagem{
		Nome:   "Matheus",
		Raca:   "Humano",
		Classe: "Mago",
		Nivel:  1,
	},
		Tipo: "matheusdcunha",
	}

	showTipo(j)
	fmt.Println(calcForcaAtaque(j))
}
