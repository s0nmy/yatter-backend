package auth

import (
	"yatter-backend-go/app/domain/object/auth"
)

// （研修用の説明）
// ui/api/auth/response.go
// 認証に関するエンドポイントのレスポンスを定義
// XxxResponse のような命名にする

// PostLoginResponse: ユーザーログインのレスポンス
type PostLoginResponse struct {
	Username string `json:"username"`
}

// toPostLoginResponse: ユーザーログインのレスポンスに変換
func toPostLoginResponse(credential *auth.AuthenticatedCredential) *PostLoginResponse {
	return &PostLoginResponse{
		Username: credential.Username(),
	}
}
