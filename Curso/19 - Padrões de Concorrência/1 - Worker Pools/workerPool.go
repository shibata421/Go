package main

import "fmt"

func main() {
	quantidadeWorkers := uint(4)
	tamanhoFibonacci := uint(45)
	tamanhosChannel := make(chan uint, tamanhoFibonacci)
	resultadosChannel := make(chan uint, tamanhoFibonacci)

	criarWorkers(quantidadeWorkers, tamanhosChannel, resultadosChannel)
	preencherTamanhosChannel(tamanhosChannel)
	escreverResultados(resultadosChannel)
}

func criarWorkers(qtdeWorkers uint, tamanhosChannel <-chan uint, resultadosChannel chan<- uint) {
	for i := uint(0); i < qtdeWorkers; i++ {
		go worker(tamanhosChannel, resultadosChannel)
	}
}

func worker(tamanhosChannel <-chan uint, resultadosChannel chan<- uint) {
	for tamanho := range tamanhosChannel {
		resultadosChannel <- fibonacci(tamanho)
	}

	close(resultadosChannel)
}

func fibonacci(tamanho uint) uint {
	if tamanho <= 1 {
		return 1
	}

	return fibonacci(tamanho-1) + fibonacci(tamanho-2)
}

func preencherTamanhosChannel(tamanhosChannel chan<- uint) {
	for i := uint(0); i < uint(cap(tamanhosChannel)); i++ {
		tamanhosChannel <- i
	}

	close(tamanhosChannel)
}

func escreverResultados(resultadosChannel <-chan uint) {
	for resultado := range resultadosChannel {
		fmt.Println(resultado)
	}
}
