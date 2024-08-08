package main

import (
	errorhandlers "be/errorHandlers"
	"be/models"
	"be/routers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	defer errorhandlers.NormalError()
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	go models.InitialTables()
	r := mux.NewRouter()
	routers.Routing(r)
	log.Println("Listening on 8080")
	http.ListenAndServe(":8080", r)
}
