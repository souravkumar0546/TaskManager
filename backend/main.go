package main

import (
	"log"
	"net/http"
	"task-manager-backend/routers"
	"task-manager-backend/utils"

	"github.com/joho/godotenv"
)

const (
	exitOK int = iota
	exitError
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	utils.InitDB()
	router := routers.InitRouter()
	//log.Println("Listening on Port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

