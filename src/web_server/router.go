package main

import (
	"net/http"
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
	handler, exist := r.FindHandler(request.URL.Path)
	if !exist {
		writer.WriteHeader(http.StatusNotFound)
		return // return para salirnos de la funcion
	}
	handler(writer, request)
}

func (r *Router) FindHandler(path string) (http.HandlerFunc, bool) {
	handler, exist := r.reglas[path]
	return handler, exist
}
