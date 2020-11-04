package main

import (
	"fmt"
	"math"
)

func main() {
	figuras := make([]area, 0)
	figuras = append(figuras, &cuadrado{lado: 4.5})
	figuras = append(figuras, &triangulo{base: 2.25, altura: 10})
	figuras = append(figuras, &circulo{radio: 2.8})
	for _, f := range figuras {
		f.calcularArea()
	}
}

type area interface {
	calcularArea() 
}

type cuadrado struct {
	lado float64
}

func (c *cuadrado) calcularArea()  {
	fmt.Println("Cuadrado, Area:", math.Pow(c.lado, 2))
} 

type circulo struct {
	radio float64
}

func (c *circulo) calcularArea() {
	fmt.Println("Circulo, Area:", math.Pi * math.Pow(c.radio, 2))
}

type triangulo struct {
	base, altura float64
}

func (t *triangulo) calcularArea() {
	fmt.Println("Triangulo, Area:", (t.base * t.altura) / 2)
}


