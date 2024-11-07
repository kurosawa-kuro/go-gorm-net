package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	AppEnv      string
}

func LoadConfig() *Config {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	// プロジェクトのルートディレクトリを取得
	currentDir, err := os.Getwd()
	if err != nil {
		log.Printf("Warning: カレントディレクトリの取得に失敗しました: %v", err)
	}

	// testsディレクトリから実行された場合のパス解決
	if filepath.Base(currentDir) == "integration" {
		currentDir = filepath.Join(currentDir, "..", "..")
	}

	// 環境に応じた.envファイルを読み込む
	envFile := filepath.Join(currentDir, fmt.Sprintf(".env.%s", env))
	if err := godotenv.Load(envFile); err != nil {
		log.Printf("Warning: %sファイルが見つかりません。環境変数を直接使用します。\n", envFile)
	}

	// 設定値をログ出力（デバッグ用）
	dbURL := os.Getenv("DATABASE_URL")
	log.Printf("Database URL: %s", dbURL)

	return &Config{
		DatabaseURL: dbURL,
		AppEnv:      os.Getenv("APP_ENV"),
	}
}
