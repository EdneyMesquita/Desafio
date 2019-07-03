package controllers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type (
	Error struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}

	Login struct {
		ID         int    `json:"id"`
		UUIDUser   string `json:"uuiduser"`
		AvatarURL  string `json:"avatarurl"`
		AvatarType string `json:"avatartype"`
		Name       string `json:"name"`
		DataStart  string `json:"datastart"`
	}

	Success struct {
		Status   string      `json:"status"`
		Message  string      `json:"message"`
		TokenJWT interface{} `json:"tokenjwt"`
		Expires  string      `json:"expires"`
		TokenMsg string      `json:"tokenmsg"`
		Login    Login       `json:"Login"`
	}
)

const AUTHTYPE = "Basic"

func Auth(writer http.ResponseWriter, request *http.Request) {
	var httpCode int
	var stringToReturn string

	if request.Header.Get("Authorization") != "" {
		writer.WriteHeader(http.StatusOK)
		requestToken := (strings.Split(request.Header.Get("Authorization"), AUTHTYPE))[1]

		credentials, err := base64.StdEncoding.DecodeString(strings.TrimSpace(requestToken))
		if err != nil {
			fmt.Fprintf(writer, err.Error())
		}

		userData := strings.Split(string(credentials), ":")

		//Consultar banco de dados com os dados contidos em userData

		httpCode = http.StatusOK
		stringToReturn = fmt.Sprintf("POST - %v", userData)
	} else {
		error := &Error{
			Status:  "error",
			Message: "Usuário não pôde ser autenticado!",
		}
		jsonString, _ := json.Marshal(error)

		httpCode = http.StatusInternalServerError
		stringToReturn = string(jsonString)
	}

	writer.WriteHeader(httpCode)
	fmt.Fprintf(writer, stringToReturn)
}
