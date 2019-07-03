package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func User(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	writer.WriteHeader(http.StatusOK)

	fmt.Fprintf(writer, "User Controller\n\n %v", params)
}
