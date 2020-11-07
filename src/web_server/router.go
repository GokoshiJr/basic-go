package main

import (
	"net/http"
)

type Router struct {
	reglas map[string]map[string]http.HandlerFunc
}

func NewRouter() *Router{
	return &Router {
		reglas: make(map[string]map[string]http.HandlerFunc),
	}
}

func (r *Router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	handler, methodExist, pathExist := r.FindHandler(request.URL.Path, request.Method)
	if !pathExist {
		writer.WriteHeader(http.StatusNotFound)
		return // return para salirnos de la funcion
	}
	if !methodExist {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	handler(writer, request)
}

func (r *Router) FindHandler(path string, method string) (http.HandlerFunc, bool, bool) {
	_, pathExist := r.reglas[path]
	handler, methodExist := r.reglas[path][method]
	return handler, methodExist, pathExist
}
