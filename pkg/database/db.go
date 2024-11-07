package database

import (
	"log"

	"go-gorm-net/internal/models"
	"go-gorm-net/pkg/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Initialize(cfg *config.Config) {
	var err error
	DB, err = gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// マイグレーション
	DB.AutoMigrate(&models.Micropost{})
}
