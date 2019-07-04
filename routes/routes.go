package routes

import (
	"Desafio/controllers"
	"Desafio/utils"
	"encoding/json"
	"fmt"
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
	r.HandleFunc("/user/{uuid}", controllers.RemoveUser).Methods("DELETE")

	r.Use(loggingMiddleware)
	http.ListenAndServe(":8080", r)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("tokenjwt")
		uuid := r.Header.Get("uuid")

		if r.RequestURI == "/auth" {
			next.ServeHTTP(w, r)
		} else {
			if token != "" {
				_, validToken := utils.DecodeToken(token, []byte(uuid))
				if validToken != false {
					next.ServeHTTP(w, r)
				} else {
					throwJSONError(w, "Token inválido!")
				}
			} else {
				throwJSONError(w, "Token inválido!")
			}
		}
	})
}

func throwJSONError(w http.ResponseWriter, message string) {
	jsonmsg := utils.JSONMsg{
		Status:  "error",
		Message: message,
	}
	jsonString, _ := json.Marshal(jsonmsg)

	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, string(jsonString))
}
