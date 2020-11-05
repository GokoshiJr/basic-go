package main

import( 
	"net/http"
	"fmt"
	"io"
)

func main() {
	req, err := http.Get("http://google.com") // http, protocolo de comunicacion usado en la web 
	if err != nil {
		fmt.Println(err)	
	} else {
		escritor := escritorWeb{}
		io.Copy(escritor, req.Body)
	}
}

type escritorWeb struct {}

// Implementa la interfaz Writter
func (escritorWeb) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	return len(p), nil
}