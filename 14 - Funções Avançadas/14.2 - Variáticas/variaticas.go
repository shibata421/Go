package main

import "fmt"

func soma(numeros ...int) int {
	var soma int

	for _, numero := range numeros {
		soma += numero
	}

	return soma
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	fmt.Println(soma(1, 2, 3))
	escrever("Oi", 1, 2, 3)
}
