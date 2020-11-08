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
	fmt.Fprintf(writer, "Payload %v\n", metaData)
}

func UserPostRequest(writer http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Fprintf(writer, "Error: %v", err)
		return
	}
	response, err1 := user.toJSON()
	if err1 != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(response) // Write espera el slice de bytes
}
