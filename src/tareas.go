package main

import (
	"fmt"
)

func main() {
	t := task {
		nombre: "Completar mi curso de Go",
		descripcion: "Haciendolo en Platzi",
	}
	// +v significa formatear como una interface clave-valor
	fmt.Printf("%+v\n", t)
	t.marcarCompleta()
	t.actualizarNombre("Finalizar mi curso de Go")
	t.actualizarDescripcion("Completar mi curso temprano")
	fmt.Printf("%+v\n", t)
}

/* Receivers de tipo task */
func (t task) marcarCompleta() {
	t.completado = true
}

func (t task) actualizarDescripcion(descripcion string) {
	t.descripcion = descripcion
}

func (t task) actualizarNombre(nombre string) {
	t.nombre = nombre
}

type task struct {
	nombre string
	descripcion string
	completado bool
}