package controllers

import (
	"Desafio/database"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type (
	//Error is the struct that will be converted to JSON referring to error 500
	Error struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}

	//Login is the struct that maps the values received from the database
	Login struct {
		ID         int    `json:"id"`
		UUIDUser   string `json:"uuiduser"`
		AvatarURL  string `json:"avatarurl"`
		AvatarType string `json:"avatartype"`
		Name       string `json:"name"`
		DataStart  string `json:"datastart"`
	}

	//ResultData is the struct that will be converted into the JSON of the response if the code is 200
	ResultData struct {
		Status   string      `json:"status"`
		Message  string      `json:"message"`
		TokenJWT interface{} `json:"tokenjwt"`
		Expires  string      `json:"expires"`
		TokenMsg string      `json:"tokenmsg"`
		Login    Login       `json:"Login"`
	}
)

//AUTHTYPE is the type of Authorization
const AUTHTYPE = "Basic"

//Auth is the function responsible for performing user verification and validation
func Auth(writer http.ResponseWriter, request *http.Request) {
	if request.Header.Get("Authorization") != "" {
		writer.WriteHeader(http.StatusOK)
		requestToken := (strings.Split(request.Header.Get("Authorization"), AUTHTYPE))[1]

		credentials, err := base64.StdEncoding.DecodeString(strings.TrimSpace(requestToken))
		if err != nil {
			fmt.Fprintf(writer, err.Error())
		}

		userData := strings.Split(string(credentials), ":")

		conn := database.Connect()

		sql := fmt.Sprintf("SELECT "+
			"u.id, u.uuiduser, a.url, a.type, u.name, u.datastart "+
			"FROM users u "+
			"LEFT JOIN avatar a ON(u.avatar = a.id) "+
			"WHERE u.email = '%s' AND u.password = '%s'", userData[0], userData[1])

		rows, _ := conn.Query(sql)
		defer rows.Close()
		conn.Close()

		var login Login
		for rows.Next() {
			rows.Scan(&login.ID, &login.UUIDUser, &login.AvatarURL, &login.AvatarType, &login.Name, &login.DataStart)
		}

		if login.ID == 0 {
			throwJSONError(writer, "Usuário não pôde ser autenticado!")
			return
		}

		resultdata := ResultData{
			Status:   "success",
			Message:  "Usuário encontrado e token gerado",
			TokenJWT: "",
			Expires:  "",
			TokenMsg: "use o token para acessar os endpoints!",
			Login:    login,
		}

		writer.WriteHeader(http.StatusOK)
		stringToReturn, _ := json.Marshal(resultdata)
		fmt.Fprintf(writer, string(stringToReturn))
	} else {
		throwJSONError(writer, "Usuário não pôde ser autenticado!")
	}
}

func throwJSONError(writer http.ResponseWriter, message string) {
	error := &Error{
		Status:  "error",
		Message: message,
	}
	jsonString, _ := json.Marshal(error)

	writer.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(writer, string(jsonString))
}
