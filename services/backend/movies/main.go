package main

import (
	"log"
	"moovio/libs/helper"
	. "moovio/services/backend/movies/controller"

	"github.com/joho/godotenv"
)

func main() {
	log.Println("Reading config...")
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("can't find .env local")
	}

	log.Println("Config loaded...")
	log.Println("Initiate DB Connection...")
	db, err := helper.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB Connected...")

	log.Println("Starting Movies Services...")
	svc := NewMovieService(db)

	apiserver := NewMoviesApiServer(svc)

	log.Fatal(apiserver.Start(":9003"))
}
