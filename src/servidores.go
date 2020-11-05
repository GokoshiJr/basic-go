package main

import (
	"net/http"
	"fmt"
	"time"
)

// Ejemplo de concurrencia = procesos asincronos

func main() {

	inicio := time.Now()

	servidores := []string{
		"http://platzi.com",
		"http://google.com",
		"http://twitter.com",
		"http://github.com",
	}
	for _, servidor := range servidores {
		revisarServidor(servidor)
	}
	tiempoTotal := time.Since(inicio)
	fmt.Println("Tiempo transcurrido:", tiempoTotal)
}

func revisarServidor(servidor string) {
	_, err := http.Get(servidor)
	if err != nil {
		fmt.Println(servidor, "no esta disponible")
	} else {
		fmt.Println(servidor, "esta funcionando")
	}
}