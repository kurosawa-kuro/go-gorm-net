package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"go-gorm-net/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Micropost モデルの定義
type Micropost struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}

var db *gorm.DB

func main() {
	// 設定を読み込む
	cfg := config.LoadConfig()

	// データベース接続
	var err error
	db, err = gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// マイグレーション
	db.AutoMigrate(&Micropost{})

	// ルーティング設定
	http.HandleFunc("/microposts", handleMicroposts)
	http.HandleFunc("/microposts/", handleMicropost)

	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleMicroposts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// 全件取得
		var microposts []Micropost
		db.Find(&microposts)
		json.NewEncoder(w).Encode(microposts)

	case http.MethodPost:
		// 新規作成
		var micropost Micropost
		json.NewDecoder(r.Body).Decode(&micropost)
		db.Create(&micropost)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(micropost)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func handleMicropost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// URLからIDを取得
	id := strings.TrimPrefix(r.URL.Path, "/microposts/")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// ID指定での取得
	var micropost Micropost
	idInt, _ := strconv.Atoi(id)
	result := db.First(&micropost, idInt)
	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(micropost)
}
