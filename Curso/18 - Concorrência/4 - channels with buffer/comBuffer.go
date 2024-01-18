package main

import "fmt"

func main() {
	canal := make(chan string, 2) // canal tem capacidade de 2

	canal <- "Olá mundo"
	canal <- "Programando em go"

	mensagem := <-canal
	fmt.Println(mensagem)

	mensagem = <-canal
	fmt.Println(mensagem)
}
