package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	config := LoadConfig()
	ConnectDB(config.DatabaseURL)
	router := SetupRoutes()
	port := os.Getenv("PORT")
	log.Println("Server running on :" + port)
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatal(err)
	}
}
