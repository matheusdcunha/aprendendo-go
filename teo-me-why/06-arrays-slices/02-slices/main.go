package main

import "fmt"

func main() {

	numeros := []int{1, 2}

	dias := []int{1, 2, 5: 6}

	sliceTamanhoPrevio := make([]int, 5, 10)

	fmt.Printf("%p\n", sliceTamanhoPrevio)
	fmt.Println(numeros)
	fmt.Println("Slice pulando indices:", dias)
	fmt.Println("Slice com tamanho prévio:", cap(sliceTamanhoPrevio))

	test := make([]int, 4)

	for i := range 6 {
		test = append(test, i)
		fmt.Printf("%p\n", test)
	}

}
