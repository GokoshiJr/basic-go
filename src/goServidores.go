package main

import (
	"net/http"
	"fmt"
	"time"
)

// Ejemplo de concurrencia = procesos asincronos

func main() {
	// channel, permiten saber que paso con las sub rutinas
	inicio := time.Now()
	// indicamos que tipo de dato se va a transmitir por los canales
	canal := make(chan string)
	servidores := []string{
		"http://platzi.com",
		"http://google.com",
		"http://twitter.com",
		"http://github.com",
	}

	// no hay while en go, hay que usar el for como si fuera el while
	i := 0
	for {
		if i > 2 {
			break
		}
		for _, servidor := range servidores {
			// go, crea un nuevo hilo/proceso para que se ejecute independientemente de main()
			go revisarServidor(servidor, canal)
		}
		time.Sleep(1 * time.Second)
		fmt.Println(<- canal)
		i++
	}

	

	// <- canal, pasa la data
	/* 
	for i:=0; i<len(servidores); i++ {
		fmt.Println(<- canal)
	} 
	*/

	tiempoTotal := time.Since(inicio)
	fmt.Println("Tiempo transcurrido:", tiempoTotal)
}

func revisarServidor(servidor string, canal chan string) {
	_, err := http.Get(servidor)
	if err != nil {
		// fmt.Println(servidor, "no esta disponible")
		canal <- servidor + " no esta disponible" // operador <- agarra lo que esta a la derecha y lo transmite
	} else {
		// fmt.Println(servidor, "esta funcionando")
		canal <- servidor + " esta funcionando"
	}
}