package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("Olá mundo")
	escrever("Programando em GO")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
