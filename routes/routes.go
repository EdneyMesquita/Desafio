package routes

import (
	"Desafio/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

//Serve is the function that defines and provides all the routes of the application
func Serve() {
	r := mux.NewRouter()

	r.HandleFunc("/ping", controllers.Ping).Methods("POST")
	r.HandleFunc("/auth", controllers.Auth).Methods("POST")

	r.HandleFunc("/user/{uuid}", controllers.User).Methods("GET")
	r.HandleFunc("/user", controllers.AddUser).Methods("POST")
	r.HandleFunc("/user", controllers.EditUser).Methods("PUT")
	r.HandleFunc("/user", controllers.RemoveUser).Methods("DELETE")

	http.ListenAndServe(":8080", r)
}
