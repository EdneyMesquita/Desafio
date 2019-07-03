package controllers

import (
	"fmt"
	"net/http"
)

func Ping(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)

	fmt.Fprintf(writer, "Endpoint PING")
}
