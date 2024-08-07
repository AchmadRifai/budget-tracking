package main

import (
	errorhandlers "be/errorHandlers"
	"be/models"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	defer errorhandlers.NormalError()
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	models.InitialTables()
	log.Println("Hello")
}
