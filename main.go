package main

import (
	"api-go-rest/database"
	"api-go-rest/routes"
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	database.ConnectDatabase()

	fmt.Println("Servidor GO Rest ouvindo em: http://localhost:8000")
	routes.HandleRequest()
}
