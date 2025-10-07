package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

func Init(){
	err := godotenv.Load()
	if err != nil {
		fmt.Println("[ERROR].env file not found, using environment variables", err)
	}

	InitDatabase()

}