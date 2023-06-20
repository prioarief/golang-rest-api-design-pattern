package configs

import (
	"os"

	"github.com/prioarief/golang-rest-api-design-pattern/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Database() *gorm.DB {
	// get env value
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUsername := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// url := "root:password@tcp(localhost:3306)/golang"
	url := dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName
	println(url)

	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		panic("Failed connect to database")
	} else {
		println("Database connected")
	}

	db.AutoMigrate(&models.Todo{})

	return db
}
