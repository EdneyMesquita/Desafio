package controllers

import (
	"Desafio/database"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

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

	Success struct {
		Status  string `json:"status"`
		Message string `json:"message"`
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
	err := request.ParseForm()
	if err != nil {
		throwJSONError(writer, "Não foi possível obter os dados!")
		return
	}

	userdata := UserData{
		AvatarURL:  request.FormValue("avatarurl"),
		AvatarType: request.FormValue("avatartype"),
		Name:       request.FormValue("name"),
		CPF:        request.FormValue("cpf"),
		Email:      request.FormValue("email"),
		Password:   request.FormValue("password"),
	}

	conn := database.Connect()
	sqlInsertAvatar := fmt.Sprintf("INSERT INTO avatar (url, type) VALUES ('%s', '%s')", userdata.AvatarURL, userdata.AvatarType)

	result, err := conn.Exec(sqlInsertAvatar)
	if err != nil {
		throwJSONError(writer, err.Error())
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		throwJSONError(writer, "Não foi possível inserir o Usuário!")
		return
	}

	var id int
	err = conn.QueryRow("SELECT id FROM avatar ORDER BY id DESC LIMIT 1").Scan(&id)

	dataStart := time.Now().Format("01-02-2006")
	sqlInsertUser := fmt.Sprintf(`INSERT INTO users (name, email, password, cpf, datastart, avatar) 
									VALUES ('%s', '%s', '%s', '%s', '%s', %d)`, userdata.Name, userdata.Email, userdata.Password, userdata.CPF, dataStart, id)

	result, err = conn.Exec(sqlInsertUser)
	if err != nil {
		throwJSONError(writer, err.Error())
		return
	}

	rowsAffected, _ = result.RowsAffected()
	if rowsAffected == 0 {
		throwJSONError(writer, "Não foi possível inserir o Usuário!")
		return
	}

	success := Success{
		Status:  "success",
		Message: "Usuário inserido com sucesso!",
	}

	jsonString, _ := json.Marshal(success)

	writer.WriteHeader(http.StatusOK)
	fmt.Fprintf(writer, string(jsonString))
}

func EditUser(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "PUT edit user")
}

func RemoveUser(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	uuid := params["uuid"]

	if uuid == "" {
		throwJSONError(writer, "Não foi possível obter os dados!")
		return
	}

	conn := database.Connect()

	var idAvatar int
	err := conn.QueryRow(fmt.Sprintf("SELECT a.id FROM users u LEFT JOIN avatar a ON(a.id = u.avatar) WHERE a.uuiduser = '%s' LIMIT 1", uuid)).Scan(&idAvatar)

	if idAvatar != 0 {
		sqlDeleteAvatar := fmt.Sprintf("DELETE FROM avatar WHERE id = %d ", idAvatar)

		result, err := conn.Exec(sqlDeleteAvatar)
		if err != nil {
			throwJSONError(writer, err.Error())
			return
		}

		rowsAffected, _ := result.RowsAffected()
		if rowsAffected == 0 {
			throwJSONError(writer, "Não foi possível excluir o Usuário!")
			return
		}
	}

	sqlDeleteUser := fmt.Sprintf("DELETE FROM users WHERE uuiduser = '%s'", uuid)

	result, err := conn.Exec(sqlDeleteUser)
	if err != nil {
		throwJSONError(writer, err.Error())
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		throwJSONError(writer, "Não foi possível excluir o Usuário!")
		return
	}

	success := Success{
		Status:  "success",
		Message: "Usuário excluído com sucesso!",
	}

	jsonString, _ := json.Marshal(success)

	writer.WriteHeader(http.StatusOK)
	fmt.Fprintf(writer, string(jsonString))
}
