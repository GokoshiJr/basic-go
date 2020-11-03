package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
	"strconv"
)

func main() {
	entrada := leerEntrada(false)
	operador := leerEntrada(true)
	// instanciar el struct
	c := calc{}
	fmt.Println("Resultado =", c.operate(entrada, operador))
}

// Lee por consola
func leerEntrada(operador bool) string {
	scanner := bufio.NewScanner(os.Stdin) // creamos un objeto Scanner
	if operador {
		fmt.Print("Ingrese un operador (+,-,*,/): ")
	} else {
		fmt.Print("Ingrese la operacion completa (ej 1+1): ")
	}
	scanner.Scan() // activamos el scanner
	return scanner.Text() // lee un string
}

// Transforma un string a un entero
func parsear(entrada string) int {
	num, _ := strconv.Atoi(entrada)
	return num
}

// (calc) receirves function, le asignamos la funcion al struct calc
func (calc) operate(entrada string, operador string) int {
	entradaLimpia := strings.Split(entrada, operador)
	num1 := parsear(entradaLimpia[0])
	num2 := parsear(entradaLimpia[1])
	switch operador {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		return num1 / num2
	default:
		return 0
	}
}

// Structs = Clases
type calc struct {}