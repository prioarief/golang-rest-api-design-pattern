package main

import (
	"github.com/prioarief/golang-rest-api-design-pattern/configs"
	"github.com/prioarief/golang-rest-api-design-pattern/routes"
)

func main() {
	db := configs.Database()
	r := routes.Router(db)

	r.Run(":8080")

}
