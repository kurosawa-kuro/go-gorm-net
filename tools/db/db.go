package db

import (
	"go-gorm-net/config"
	"go-gorm-net/database"
	"go-gorm-net/models"
	"log"
)

func ResetDB() {
	CleanupDatabase()
	SeedDatabase()
}

func CleanupDatabase() {
	cfg := config.LoadConfig()
	database.Initialize(cfg)

	database.DB.Migrator().DropTable(&models.Micropost{})
	database.DB.AutoMigrate(&models.Micropost{})

	log.Println("Database cleaned up successfully")
}

func SeedDatabase() {
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
