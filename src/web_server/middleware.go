package main

import (
	"net/http"
	"fmt"
)

func CheckAuth() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			flag := true
			fmt.Println("Checking Authentication")
			if flag {
				f(w, r)
			} else {
				return
			}
		}
	}
}