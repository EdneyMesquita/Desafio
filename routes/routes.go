package routes

import (
	"Desafio/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

//Serve is the function that defines and provides all the routes of the application
func Serve() {
	r := mux.NewRouter()

	r.HandleFunc("/auth", controllers.Auth).Methods("POST")
	r.HandleFunc("/user/{uuid}", controllers.User).Methods("GET")

	http.ListenAndServe(":8080", r)
}
