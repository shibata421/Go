package main

import (
	"fmt"
	"math"
)

type forma interface {
	area() float64
}

type retangulo struct {
	base   float64
	altura float64
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func (r retangulo) area() float64 {
	return r.altura * r.base
}

func escreverArea(f forma) {
	fmt.Printf("A área da forma é %0.2f\n", f.area())
}

func main() {
	retang := retangulo{1, 2}
	circ := circulo{3}

	escreverArea(retang)
	escreverArea(circ)
}
