package configs

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Database() *gorm.DB {
	url := "root:password@tcp(localhost:3306)/golang"

	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		panic("Failed connect to database")
	} else {
		println("Database connected")
	}

	return db
}
