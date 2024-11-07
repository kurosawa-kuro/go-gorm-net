package integration

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"go-gorm-net/config"
	"go-gorm-net/handlers"
	"go-gorm-net/models"
	"go-gorm-net/tools/db"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// テスト環境の設定
	os.Setenv("APP_ENV", "test")

	// 設定を明示的に読み込む
	cfg := config.LoadConfig()

	// 設定値の確認のためにログ出力
	log.Printf("Using database URL: %s", cfg.DatabaseURL)

	// テストの実行
	code := m.Run()

	os.Exit(code)
}

func setupTest(t *testing.T) {
	// データベースのリセットと初期データの投入
	db.ResetDB()
}

func TestGetAllMicroposts(t *testing.T) {
	// テストのセットアップ
	setupTest(t)

	// ハンドラーの準備
	handler := handlers.NewMicropostHandler()

	// テストサーバーの作成
	server := httptest.NewServer(http.HandlerFunc(handler.HandleMicroposts))
	defer server.Close()

	// リクエストの実行
	resp, err := http.Get(server.URL)
	assert.NoError(t, err)
	defer resp.Body.Close()

	// ステータスコードの検証
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// レスポンスボディの検証
	var microposts []models.Micropost
	err = json.NewDecoder(resp.Body).Decode(&microposts)
	assert.NoError(t, err)

	// 期待される結果の検証
	assert.Len(t, microposts, 3) // シードデータで3件投入されていることを想定
	assert.Equal(t, "最初の投稿", microposts[0].Title)
	assert.Equal(t, "2番目の投稿", microposts[1].Title)
	assert.Equal(t, "3番目の投稿", microposts[2].Title)
}

// POSTリクエストのテスト
func TestCreateMicropost(t *testing.T) {
	// テストのセットアップ
	setupTest(t)

	// ハンドラーの準備
	handler := handlers.NewMicropostHandler()

	// テストサーバーの作成
	server := httptest.NewServer(http.HandlerFunc(handler.HandleMicroposts))
	defer server.Close()

	// 新しい投稿データの作成
	newPost := models.Micropost{
		Title: "テスト投稿",
	}
	jsonData, err := json.Marshal(newPost)
	assert.NoError(t, err)

	// POSTリクエストの実行
	resp, err := http.Post(server.URL, "application/json", bytes.NewBuffer(jsonData))
	assert.NoError(t, err)
	defer resp.Body.Close()

	// ステータスコードの検証
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	// レスポンスボディの検証
	var createdPost models.Micropost
	err = json.NewDecoder(resp.Body).Decode(&createdPost)
	assert.NoError(t, err)

	// 作成された投稿の検証
	assert.Equal(t, "テスト投稿", createdPost.Title)
	assert.NotZero(t, createdPost.ID)
}