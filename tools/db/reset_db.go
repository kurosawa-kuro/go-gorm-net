package db

import (
	"log"

	"go-gorm-net/config"
	"go-gorm-net/database"
	"go-gorm-net/models"
)

func ResetDB() {
	CleanupDatabase()
	SeedDatabase()
}

func CleanupDatabase() {
	cfg := config.LoadConfig()
	database.Initialize(cfg)

	// テーブルを削除して再作成
	database.DB.Migrator().DropTable(&models.Micropost{})
	database.DB.AutoMigrate(&models.Micropost{})

	log.Println("Database cleaned up successfully")
}

func SeedDatabase() {
	// データベース接続は既に Initialize で行われているため不要

	// シードデータの作成
	posts := []models.Micropost{
		{Title: "最初の投稿"},
		{Title: "2番目の投稿"},
		{Title: "3番目の投稿"},
	}

	for _, post := range posts {
		database.DB.Create(&post)
	}

	log.Println("Database seeded successfully")
}
