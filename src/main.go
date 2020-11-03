package main // nombre del paquete

import "fmt" // librerias necesarias para ejecutar el programa 

// go build main.go - se crea el ejecutable.exe del programa
// go run main.go - compila, ejecuta y elimina el ejecutable.exe 

func main() {
	// decalracion de variables
	var mensaje string = "Hola mundo" // declaracion mas tipada
	mensajeFacil := "Hola mundo usando := " // para inferir el tipo
	fmt.Println(mensaje)
	fmt.Println(mensajeFacil)
	// numeros
	a := 10. // float porque termina en punto
	var b float64 = 3
	fmt.Println(a / b)
	var c int = 10
	d := 3 
	fmt.Println(c + d)
	// boolean
	x := true
	y := true
	var z bool = false
	fmt.Println(x && y && z)
	privada()
	Publica()
	imprimirHola()
}

// si la funcion empieza en minuscula es privada
func privada() {
	fmt.Println("Ejecutar logica que no necesita ser exportada")
}

// si la funcion empieza en mayuscula es publica
func Publica() {
	fmt.Println("Logica que quiero exportar")
}

// defer, indica que la linea diferida si ejecuta al final de la funcion
func imprimirHola()  {
	defer fmt.Println("Mundo")
	fmt.Println("Hola")
}
