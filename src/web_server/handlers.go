package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleRoot(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(writer, "Hola mundo")
}

func HandleHome(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(writer, "This is the API Endpoint")
}

func PostRequest(writer http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var metaData MetaData
	err := decoder.Decode(&metaData)
	if err != nil {
		fmt.Fprintf(writer, "Error: %v", err)
		return
	}
	fmt.Println(metaData["name"])
	fmt.Fprintf(writer, "Payload %v\n", metaData)
}
