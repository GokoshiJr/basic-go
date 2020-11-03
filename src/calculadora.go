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

	fmt.Print("Ingrese una operacion: ")
	scanner.Scan()
	operacion := scanner.Text()
	
	// fmt.Println(operacion)	
	valores := strings.Split(operacion, "+")
	// fmt.Println(valores)

	// strcon.Atoi lo pasa de string a enteros
	operador1, _ := strconv.Atoi(valores[0])
	operador2, _ := strconv.Atoi(valores[1])
	// fmt.Println("Resultado", valores[0] + valores[1])
	fmt.Println("Resultado:", operador1 + operador2)
}