package main

import (
	"github.com/joho/godotenv"
	"golang-rest-api/database"
	"golang-rest-api/models"
	"golang-rest-api/routes"
	"os"
)

func main() {
	database.Init()
	godotenv.Load(".env")
	port := os.Getenv("PORT")
	models.pariwisata()
	routes.Init(port)
}
