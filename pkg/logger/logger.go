package logger

import (
	"log"
	"os"
	"path/filepath"
)

var (
	AccessLogger *log.Logger
	ErrorLogger  *log.Logger
)

func Initialize() {
	// ログディレクトリの作成
	logDir := "logs"
	if err := os.MkdirAll(logDir, 0755); err != nil {
		log.Fatal(err)
	}

	// アクセスログファイルの設定
	accessLogFile, err := os.OpenFile(
		filepath.Join(logDir, "access.log"),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666,
	)
	if err != nil {
		log.Fatal(err)
	}

	// エラーログファイルの設定
	errorLogFile, err := os.OpenFile(
		filepath.Join(logDir, "error.log"),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666,
	)
	if err != nil {
		log.Fatal(err)
	}

	AccessLogger = log.New(accessLogFile, "", log.LstdFlags)
	ErrorLogger = log.New(errorLogFile, "ERROR: ", log.LstdFlags|log.Lshortfile)
}
