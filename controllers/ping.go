package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type retorno struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

//Ping returns a test JSON string
func Ping(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)

	retorno := &retorno{
		Status:  "success",
		Message: "Endpoint PING",
	}

	jsonString, _ := json.Marshal(retorno)

	fmt.Fprintf(writer, string(jsonString))
}
