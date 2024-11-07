package develop

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Micropost struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}

func CleanupDatabase() {
	dsn := "postgresql://postgres:postgres@localhost:5432/web_app_db_integration_go?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// テーブルを削除して再作成
	db.Migrator().DropTable(&Micropost{})
	db.AutoMigrate(&Micropost{})

	log.Println("Database cleaned up successfully")
}
