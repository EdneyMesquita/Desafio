package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	//DbUser is the username for database connection
	DbUser = "postgres"
	//DbPassword is the password for database connection
	DbPassword = "postgres"
	//DbName is the name of the schema that will be connected
	DbName = "desafio"
)

//Connect is responsible for opening a connection to the database
func Connect() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DbUser, DbPassword, DbName)

	db, _ := sql.Open("postgres", dbinfo)

	return db
}
