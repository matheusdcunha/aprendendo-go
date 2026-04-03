package main

import "fmt"

func addOne(a int) int {
	a++
	return a
}

func addOnePonteiro(a *int) int {
	*a++
	return *a
}

func main() {

	var x int = 42
	p := &x

	addOne(*p)
	addOnePonteiro(p)

	fmt.Println("Valor de P: ", *p)
	fmt.Println("Endereço de P: ", p)
}
