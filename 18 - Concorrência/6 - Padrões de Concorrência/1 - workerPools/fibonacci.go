package main

import "fmt"

var tamanhoFibonacci uint
var tamanhosChannel chan uint
var resultadosChannel chan uint

func init() {
	tamanhoFibonacci = uint(45)
	resultadosChannel = make(chan uint, tamanhoFibonacci)
	tamanhosChannel = make(chan uint, tamanhoFibonacci)
}

func main() {
	criarWorkers(4)
	gerarTamanhosChannel()
	escreverResultados()
}

func criarWorkers(qtdeWorkers uint) {
	for i := uint(0); i < qtdeWorkers; i++ {
		go worker()
	}
}

func worker() {
	for tamanho := range tamanhosChannel {
		resultadosChannel <- fibonacci(tamanho)
	}
}

func fibonacci(tamanho uint) uint {
	if tamanho <= 1 {
		return 1
	}

	return fibonacci(tamanho-1) + fibonacci(tamanho-2)
}

func gerarTamanhosChannel() {
	for i := uint(0); i < tamanhoFibonacci; i++ {
		tamanhosChannel <- i
	}

	close(tamanhosChannel)
}

func escreverResultados() {
	for i := uint(0); i < tamanhoFibonacci; i++ {
		fmt.Println(<-resultadosChannel)
	}

	close(resultadosChannel)
}
