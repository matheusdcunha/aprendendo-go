// Escreva uma função que recebe a e b e troque seus valores. a = 1; b=2 -> a=2; b=1

package main

import "fmt"

func troca(a, b *int) {
	*a, *b = *b, *a
}

func main() {

	a := 1
	b := 2

	troca(&a, &b)

	fmt.Println("A:", a)
	fmt.Println("B:", b)

}
