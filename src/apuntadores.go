package main

import (
	"fmt"
)
/* 
	En Go hay dos tipos de parámetros que podemos usar en nuestras funciones o métodos:

	Pass by value (por valor) : Aquí la función ara una copia de la variable que se esta pasando y modificara la copia.

	Pass by pointer/reference (por puntero) : Aquí la función obtendrá un puntero que apunta a la misma dirección donde esta el valor de la variable o struc 
*/
func main() {
	x := 25
	fmt.Println(x)
	fmt.Println(&x) // mostramos la direccion en memoria
	// cambiarValor(x) // las direcciones son diferentes, porq hace una copia de esta variable
	y := &x
	fmt.Println(*y)
}

func cambiarValor(a int) {
	fmt.Println(&a)
	a = 36
}