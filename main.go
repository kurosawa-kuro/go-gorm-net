package main

import (
	"log"
	"net/http"

	"go-gorm-net/config"
	"go-gorm-net/database"
	"go-gorm-net/handlers"
)

func main() {
	cfg := config.LoadConfig()
	database.Initialize(cfg)

	http.HandleFunc("/microposts", handlers.HandleMicroposts)
	http.HandleFunc("/microposts/", handlers.HandleMicropost)

	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
