package main

import (
	"fmt"
)

func main() {
	i := 0

	for i < 10 {
		i++
		// fmt.Println("Incrementando i")
		// time.Sleep(time.Second)
	}

	for i := 0; i < 10; i++ {
		fmt.Println("Incrementando i")
		// time.Sleep(time.Second)
	}

	nomes := [3]string{"João", "Davi", "Lucas"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for indice := range nomes {
		fmt.Println(indice)
	}

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Leonardo",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
}
