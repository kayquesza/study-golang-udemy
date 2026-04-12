package formas

import (
	"math"
)

type Forma interface {
	Area() float64
}

type Retangulo struct {
	Altura  float64
	Largura float64
}

func (r Retangulo) Area() float64 {
	return r.Altura * r.Largura
}

func (c Circulo) Area() float64 {
	return math.Pi * (c.raio * c.raio)
}

type Circulo struct {
	raio float64
}
