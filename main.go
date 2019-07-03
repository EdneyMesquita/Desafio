package main

import (
	"Desafio/routes"
	"fmt"
)

func main() {
	fmt.Println("Server started at port 8080 ...")
	routes.Serve()
}
