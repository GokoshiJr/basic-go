package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
	"strconv"
)

func main() {
	// creamos un objeto Scanner
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("\nIngrese una operacion: ")
	scanner.Scan()
	operacion := scanner.Text()

	operador := "-"
	numeros := strings.Split(operacion, operador)
	num1, err1 := strconv.Atoi(numeros[0])
	num2, err2 := strconv.Atoi(numeros[1])

	// nil = null
	if (err1 != nil ) { 
		fmt.Println("Error1", err1) 
	} else if (err2 != nil) {
		fmt.Println("Error2", err2) 
	} else {
		switch operador {
			case "+":
				fmt.Println("Resultado:", num1 + num2)
			case "-":
				fmt.Println("Resultado:", num1 - num2)
			case "*":
				fmt.Println("Resultado:", num1 * num2)
			case "/":
				fmt.Println("Resultado:", num1 / num2)
			default:
				fmt.Println(operador, "No soportado")
		}
	}
}