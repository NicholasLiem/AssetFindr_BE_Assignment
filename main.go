package main

import (
	"github.com/NicholasLiem/GoLang_Microservice/database"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.DB = database.SetupDB()

	port := os.Getenv("PORT")

	router := NewRouter()

	log.Println("Running the server on port " + port)

	log.Fatal(http.ListenAndServe("127.0.0.1:"+port, router))
}
