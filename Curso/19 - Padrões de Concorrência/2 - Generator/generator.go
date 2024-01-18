package main

import (
	"fmt"
	"time"
)

func main() {
	escreverChannel := escrever("Oi")

	for texto := range escreverChannel {
		fmt.Println(texto)
	}
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			canal <- fmt.Sprintf("Texto recebido: %s, na iteração %d", texto, i)
			time.Sleep(time.Millisecond * 500)
		}

		close(canal)
	}()

	return canal
}
