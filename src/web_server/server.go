package main

import (
	"net/http"
)

type Server struct {
	port string
	router *Router
}

func NewServer(port string) *Server{
	return &Server {
		port: port,
		router: NewRouter(),
	}
}

func (servidor *Server) Listen() error {
	// puerto, handler
	http.Handle("/", servidor.router) // raiz de nuestro servidor o punto de entrada
	err := http.ListenAndServe(servidor.port, nil)
	if err != nil {
		return err
	}
	return nil
}

func (servidor *Server) Handle(path string, method string, handler http.HandlerFunc) {
	_, exist := servidor.router.reglas[path]
	if !exist {
		servidor.router.reglas[path] = make(map[string]http.HandlerFunc)
	}
	servidor.router.reglas[path][method] = handler 
}

// si agregamos los ... indicamos que no sabemos la cantidad de parametros que van a llegar
func (servidor *Server) AddMiddleware(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, middleware := range middlewares {
		f = middleware(f)
	}
	return f
}
