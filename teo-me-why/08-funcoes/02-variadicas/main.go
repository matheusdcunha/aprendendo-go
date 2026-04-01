package main

import "fmt"

func somaDois(a, b int) int {
	return a + b
}

func somaTudo(valores ...int) int {
	total := 0

	for _, v := range valores {
		total += v
	}

	return total
}

func fib(n uint) int {

	if n <= 1 {
		return int(n)
	} else {
		return fib(n-1) + fib(n-2)
	}
}

func main() {

	dois := []int{1, 2}

	fmt.Println(somaTudo(dois...))
}
