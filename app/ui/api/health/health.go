package health

import (
	"log/slog"
	"net/http"
)

// （研修用の説明）
// ui/api/health/health.go
// ヘルスチェックの実装にはテストなどを予定していないため、関数のみを定義している

func Check(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	_, err := w.Write([]byte("OK"))
	if err != nil {
		// レスポンスボディの書き込みに失敗している かつ ステータスコードはレスポンスボディ書き込み後に変更できないのでログにエラーを出す
		slog.Error("health check: failed to write response body", "err", err)
	}
}
