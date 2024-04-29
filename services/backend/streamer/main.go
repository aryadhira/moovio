package main

import (
	"log"
	"moovio/libs/helper"

	. "moovio/services/backend/streamer/controller"

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

	log.Println("Starting Streamer Services...")
	svc := NewStreamerService(db)

	apiserver := NewStreamerApiServer(svc)

	log.Fatal(apiserver.Start(":9004"))
}
