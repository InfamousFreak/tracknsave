package database 

import (
	"fmt"
	"strconv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"tracknsave/config"
)

var DB *gorm.DB

func Connect(){ //func connects to database 

	var err error 
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		fmt.Println("Didnt work")
	}

	//connection url to connect to postgres database
 	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME")) 
	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	log.Println("Database Connection established")
}