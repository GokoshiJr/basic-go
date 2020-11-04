package main

import("fmt")

func main() {
	perro1 := perro{}
	moverAnimal(perro1)
	pez1 := pez{}
	moverAnimal(pez1)
	pajaro1 := pajaro{}
	moverAnimal(pajaro1)
}

type animal interface {
	// los structs que implementan el metodo mover, 
	// implementan la interfaz animal o tipo animal 
	// sin usar palabras claves
	mover() string
}

func moverAnimal(a animal) {
	fmt.Println(a.mover())
}

type perro struct {}

func (perro) mover() string {
	return "Soy un perro y camino"
}

type pez struct {}

func (pez) mover() string {
	return "Soy un pez y estoy nadando"
}

type pajaro struct {}

func (pajaro) mover() string {
	return "Soy un pajaro y estoy volando"
}
