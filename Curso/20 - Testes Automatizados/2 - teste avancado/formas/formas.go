package formas

import (
	"math"
)

type Forma interface {
	area() float64
}

type Retangulo struct {
	Base   float64
	Altura float64
}

type Circulo struct {
	Raio float64
}

func (c Circulo) Area() float64 {
	return math.Pi * math.Pow(c.Raio, 2)
}

func (r Retangulo) Area() float64 {
	return r.Altura * r.Base
}
