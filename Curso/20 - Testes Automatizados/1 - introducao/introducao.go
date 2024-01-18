package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	tipoDeEndereco := enderecos.TipoDeEndereco("Rua dos sabiass")

	fmt.Println(tipoDeEndereco)
}
