package main

import (
	"github.com/NicholasLiem/AssetFindr_BackendAssignment/adapter"
	"github.com/NicholasLiem/AssetFindr_BackendAssignment/internal/app"
	"github.com/NicholasLiem/AssetFindr_BackendAssignment/internal/datastruct"
	"github.com/NicholasLiem/AssetFindr_BackendAssignment/internal/repository"
	"github.com/NicholasLiem/AssetFindr_BackendAssignment/internal/service"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

func main() {
	/**
	Load env file
	*/
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	/**
	DB setup
	*/
	db := repository.SetupDB()

	/**
	Registering DAO's and Services
	*/
	dao := repository.NewDAO(db)

	postService := service.NewPostService(dao)

	/**
	Registering Services to Server
	*/
	server := app.NewMicroservice(
		postService,
	)

	/**
	DB Migration
	*/
	datastruct.Migrate(db, &datastruct.Post{})
	serverRouter := adapter.NewRouter(*server)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "OPTIONS", "DELETE"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
		ExposedHeaders:   []string{},
	})

	handler := c.Handler(serverRouter)

	port := os.Getenv("BE_PORT")
	log.Println("[Server] Running the server on port " + port)

	if os.Getenv("ENVIRONMENT") == "DEV" {
		log.Fatal(http.ListenAndServe("127.0.0.1:"+port, handler))
	} else {
		log.Fatal(http.ListenAndServe(":"+port, handler))
	}
}
