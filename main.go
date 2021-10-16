package main

import (
	"github.com/joho/godotenv"
	"golang-rest-api/application/database"
	"golang-rest-api/application/models"
	"golang-rest-api/application/routes"
	"os"
)

func main() {
	db.Init()
	godotenv.Load(".env")
	port := os.Getenv("PORT")
	routes.Init(port)
}
