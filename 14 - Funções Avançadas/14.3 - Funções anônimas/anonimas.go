package main

import "fmt"

func main() {
	texto := func(texto string) string {
		return texto
	}("test")

	fmt.Println(texto)
}
