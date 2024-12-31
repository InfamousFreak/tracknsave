package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const Secret = "secret"




//config func to get env value
func Config(key string) string {
	
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	//return te value of the variable
	return os.Getenv(key)
}