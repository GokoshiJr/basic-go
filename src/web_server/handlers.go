package main

import (
	"fmt"
	"net/http"
)

func HandleRoot(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(writer, "Hola mundo")
}