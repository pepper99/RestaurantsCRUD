package main

import (
	"log"
	"swarch/poptum/restaurants/config"
	"swarch/poptum/restaurants/routes"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	mongoClient := config.DBConnect()

	routes.SetupRoutes(mongoClient)
}
