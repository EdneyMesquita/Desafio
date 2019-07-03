package controllers

import (
	"Desafio/database"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type (
	UserData struct {
		AvatarURL  string `json:"avatarurl"`
		AvatarType string `json:"avatartype"`
		Name       string `json:"name"`
		CPF        string `json:"cpf"`
		Email      string `json:"email"`
		Password   string `json:"password"`
	}

	ResponseData struct {
		UUID  string `json:"uuid"`
		Name  string `json:"name"`
		CPF   string `json:"cpf"`
		Email string `json:"email"`
	}
)

func User(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	if params["uuid"] == "" {
		error := &Error{
			Status:  "error",
			Message: "UUID não informado!",
		}
		jsonString, _ := json.Marshal(error)

		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(writer, string(jsonString))
		return
	}

	conn := database.Connect()
	sql := fmt.Sprintf("SELECT uuiduser, name, cpf, email FROM users WHERE uuiduser = '%s'", params["uuid"])

	rows, _ := conn.Query(sql)
	defer rows.Close()
	conn.Close()

	var responsedata ResponseData
	for rows.Next() {
		rows.Scan(&responsedata.UUID, &responsedata.Name, &responsedata.CPF, &responsedata.Email)
	}

	if responsedata.UUID == "" {
		throwJSONError(writer, "Usuário não encontrado!")
		return
	}

	jsonString, _ := json.Marshal(responsedata)
	writer.WriteHeader(http.StatusOK)
	fmt.Fprintf(writer, "%v\n", string(jsonString))
}

func AddUser(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	writer.WriteHeader(http.StatusOK)

	// decoder := json.NewDecoder(request.Body)

	// var responsedata ResponseData
	// decoder.Decode(&responsedata)

	fmt.Fprintf(writer, "Post\n\n%v\n", params)
}

func EditUser(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "PUT edit user")
}

func RemoveUser(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "DELETE remove user")
}
