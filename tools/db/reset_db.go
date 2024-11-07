package main

import (
	"log"

	"go-gorm-net/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ResetDB() {
	CleanupDatabase()
	SeedDatabase()
}

type Micropost struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}

func CleanupDatabase() {
	cfg := config.LoadConfig()
	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// テーブルを削除して再作成
	db.Migrator().DropTable(&Micropost{})
	db.AutoMigrate(&Micropost{})

	log.Println("Database cleaned up successfully")
}

func SeedDatabase() {
	cfg := config.LoadConfig()
	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// シードデータの作成
	posts := []Micropost{
		{Title: "最初の投稿"},
		{Title: "2番目の投稿"},
		{Title: "3番目の投稿"},
	}

	for _, post := range posts {
		db.Create(&post)
	}

	log.Println("Database seeded successfully")
}

func main() {
	ResetDB()
}
