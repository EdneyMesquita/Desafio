package utils

import (
	"fmt"
	"log"

	jose "github.com/dvsekhvalnov/jose2go"
)

type (
	JSONMsg struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}
)

func GenerateToken(payload string, key []byte) string {
	token, err := jose.Sign(payload, jose.HS256, key)
	if err != nil {
		log.Fatal(err.Error())
	}

	return token
}

func DecodeToken(token string, key []byte) {
	payload, headers, err := jose.Decode(token, key)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("%s\n\n%v", payload, headers)
}
