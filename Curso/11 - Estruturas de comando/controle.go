package main

import "fmt"

func main() {
	fmt.Println("Estrutura de controle")

	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Maior que zero")
	} else if outroNumero < 0 {
		fmt.Println("Menor que zero")
	} else {
		fmt.Println("Igual a zero")
	}

}
