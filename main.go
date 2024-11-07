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

	micropostHandler := handlers.NewMicropostHandler()
	http.HandleFunc("/microposts", micropostHandler.HandleMicroposts)
	http.HandleFunc("/microposts/", micropostHandler.HandleMicropost)

	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
