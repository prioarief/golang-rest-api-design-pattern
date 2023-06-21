package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/prioarief/golang-rest-api-design-pattern/routes"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic("Failed to load .env file")
	}

	r := routes.Router()

	appPort := os.Getenv("APP_PORT")

	r.Run(":" + appPort)

}
