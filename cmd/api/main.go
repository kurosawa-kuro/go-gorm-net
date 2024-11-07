package main

import (
	"log"
	"net/http"

	"go-gorm-net/internal/handlers"
	"go-gorm-net/internal/middleware"
	"go-gorm-net/internal/services"
	"go-gorm-net/pkg/config"
	"go-gorm-net/pkg/database"
	"go-gorm-net/pkg/logger"
)

func main() {
	// ロガーの初期化
	logger.Initialize()

	cfg := config.LoadConfig()
	database.Initialize(cfg)

	// サービスのインスタンスを作成
	micropostService := services.NewMicropostService()

	// ハンドラーにサービスを注入
	micropostHandler := handlers.NewMicropostHandler(micropostService)

	// ミドルウェアを適用したルーティング
	http.HandleFunc("/microposts", middleware.LoggingMiddleware(micropostHandler.HandleMicroposts))
	http.HandleFunc("/microposts/", middleware.LoggingMiddleware(micropostHandler.HandleMicropost))

	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
