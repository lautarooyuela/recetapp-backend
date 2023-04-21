package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnection() {
	var error error
	var host = os.Getenv("HOST")
	var user = os.Getenv("USER")
	var password = os.Getenv("PASS")
	var dbname = os.Getenv("DBNAME")
	var port = os.Getenv("DBPORT")
	var DSN = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, password, dbname, port)

	DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("Database connection successful")
	}
}
