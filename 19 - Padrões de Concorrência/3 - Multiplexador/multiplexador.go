package main

import (
	"fmt"
	"time"
)

func main() {
	c := multiplexar(escrever("Oi!", 500), escrever("Tchau!", 1000))

	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
}

func multiplexar(canalDeEntrada ...<-chan string) <-chan string {
	canalDeSaida := make(chan string)

	go func() {
		for {
			select {
			case texto := <-canalDeEntrada[0]:
				canalDeSaida <- texto
			case texto := <-canalDeEntrada[1]:
				canalDeSaida <- texto
			}
		}
	}()

	return canalDeSaida
}

func escrever(texto string, duracao int) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Texto recebido: %s", texto)
			time.Sleep(time.Millisecond * time.Duration(duracao))
		}

		close(canal)
	}()

	return canal
}
