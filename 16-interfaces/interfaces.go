package main

import (
	"fmt"
	"math"
)

type retangulo struct {
	altura  float64
	largura float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

func (c circulo) area() float64 {
	return math.Pi * (c.raio * c.raio)
}

type circulo struct {
	raio float64
}

// Criação da 'interface'
// Só podem ser passadas 'assinaturas de métodos'
type forma interface {
	area() float64
}

func escreverArea(f forma) {
	fmt.Printf("A área de forma é %0.2f.\n\n", f.area())
}

func main() {
	r := retangulo{20, 25}
	escreverArea(r)

	c := circulo{38}
	escreverArea(c)

}
