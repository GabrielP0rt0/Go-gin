package main

import (
	"go-gin/database"
	"go-gin/models"
	"go-gin/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Panic("Erro ao carregar arquivo .env: ", err)
		return
	}
	models.Config.Host = os.Getenv("DATABASE_HOST")
	models.Config.DbName = os.Getenv("DATABASE_NAME")
	models.Config.Password = os.Getenv("DATABASE_PASSWORD")
	models.Config.User = os.Getenv("DATABASE_NAME")
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
