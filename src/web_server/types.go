package main

import ("net/http")

type Middleware func(http.HandlerFunc) http.HandlerFunc

type User struct {
	Nombre string
	Email string
	Phone string
}

type MetaData interface {}
