package main

import (
	"net/http"
	"fmt"
)

type Router struct {
	reglas map[string]http.HandlerFunc
}

func NewRouter() *Router{
	return &Router {
		reglas: make(map[string]http.HandlerFunc),
	}
}

func (r *Router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// Fprintf, debemos pasarle un escritor y un mensaje
	fmt.Fprintf(writer, "Hello World, This is my server in Go")
}

func (r *Router) FindHandler(path string) (http.HandlerFunc, bool) {
	handler, exist := r.reglas[path]
	return handler, exist
}