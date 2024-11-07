package middleware

import (
	"go-gorm-net/pkg/logger"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// リクエストを処理
		next(w, r)

		// アクセスログを記録
		logger.AccessLogger.Printf(
			"method=%s path=%s remote_addr=%s duration=%v",
			r.Method,
			r.URL.Path,
			r.RemoteAddr,
			time.Since(start),
		)
	}
}
