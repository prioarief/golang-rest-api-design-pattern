package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/prioarief/golang-rest-api-design-pattern/configs"
	"github.com/prioarief/golang-rest-api-design-pattern/routes"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic("Failed to load .env file")
	}

	db := configs.Database()
	r := routes.Router(db)

	appPort := os.Getenv("APP_PORT")

	r.Run(":" + appPort)

}
