package main

import (
	"net/http"
 	"encoding/json"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

type User struct {
	Nombre string `json:"name"`
	Email string	`json:"email"`
	Phone string	`json:"phone"`
}

func (user *User) toJSON() ([]byte, error) {
	return json.Marshal(user)
}

type MetaData interface {}
