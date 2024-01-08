package main

import "fmt"

func fibonacci(tamanho uint) uint {
	if tamanho <= 1 {
		return 1
	}

	return fibonacci(tamanho-1) + fibonacci(tamanho-2)
}

func main() {
	for i := uint(0); i < 10; i++ {
		fmt.Println(fibonacci(i))
	}
}
