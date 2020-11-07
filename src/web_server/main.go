package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := NewServer(":3000")
	server.Listen()
}