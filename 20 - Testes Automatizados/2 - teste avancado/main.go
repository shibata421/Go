package main

import (
	"fmt"
	"teste-avancado/formas"
)

func main() {
	ret := formas.Retangulo{Base: 1, Altura: 2}
	fmt.Println(ret.Area())

	cir := formas.Circulo{Raio: 1}
	fmt.Println(cir.Area())
}
