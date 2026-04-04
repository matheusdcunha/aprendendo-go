package main

import "fmt"

type Motor struct {
	Potencia int
}

func (m *Motor) Tunar() {
	m.Potencia = m.Potencia * 2
}

type Carro struct {
	*Motor
	Modelo string
}

type Trator struct {
	Motor
	Ano uint
}

func main() {

	m := Motor{
		Potencia: 2,
	}

	t := Trator{
		Motor: m,
	}

	c := Carro{
		Motor: &m,
	}

	fmt.Println("Motor do Trator:", t.Potencia)
	fmt.Println("Motor do Carro:", c.Potencia)

	fmt.Println("---")
	fmt.Println("Tunando o Moto...")
	m.Tunar()
	fmt.Println("---")

	fmt.Println("Motor do Trator:", t.Potencia)
	fmt.Println("Motor do Carro:", c.Potencia)
}
