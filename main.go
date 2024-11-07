package main

import (
	"log"
	"net/http"

	"go-gorm-net/config"
	"go-gorm-net/database"
	"go-gorm-net/handlers"
	"go-gorm-net/logger"
	"go-gorm-net/middleware"
)

func main() {
	// ロガーの初期化
	logger.Initialize()

	cfg := config.LoadConfig()
	database.Initialize(cfg)

	micropostHandler := handlers.NewMicropostHandler()

	// ミドルウェアを適用したルーティング
	http.HandleFunc("/microposts", middleware.LoggingMiddleware(micropostHandler.HandleMicroposts))
	http.HandleFunc("/microposts/", middleware.LoggingMiddleware(micropostHandler.HandleMicropost))

	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
