package config

import (
	"log"
	"os"
)

func InitLogger() *os.File {

	file, err := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Impossible d'ouvrir le fichier de log : %v", err)
	}

	log.SetOutput(file)

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	return file

}
