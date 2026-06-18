package main

import (
	"log"
	"net/http"
)

func main() {
	config := LoadConfig()
	ConnectDB(config.DatabaseURL)
	router := SetupRoutes()
	log.Println("Server running on :8000")
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatal(err)
	}
}
