package main

import("fmt")

// Maps, son como los diccionarios en python o los objetos en js
// Es una estructura clave-valor
// Make permite crear mapas

func main() {
	// mapeamos una clave string a un valor entero
	mapa1 := make(map[string]int)
	mapa1["a"] = 8
	fmt.Println(mapa1)
	fmt.Println(mapa1["a"])
}