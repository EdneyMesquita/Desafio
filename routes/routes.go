package routes

import (
	"desafio/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func Serve() {
	r := mux.NewRouter()

	r.HandleFunc("/auth", controllers.Auth).Methods("POST")

	r.HandleFunc("/user/{uuid}", controllers.User).Methods("GET")
	// http.Handle("/", r)
	http.ListenAndServe(":8080", r)
}
