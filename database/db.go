package database 

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Connect(){

	dsn := "host=localhost user=postgres password=smarak dbname=exptracker port=5432 sslmode=disable"

	var err error 

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	log.Println("Database Connection established")
}