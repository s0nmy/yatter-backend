package user

import "yatter-backend-go/app/domain/object/user"

// （研修用の説明）
// ui/api/user/response.go
// ユーザーに関するエンドポイントのレスポンスを定義
// XxxResponse のような命名にする

// PostUsersResponse: ユーザー新規登録のレスポンス
type PostUsersResponse struct {
	ID             uint64 `json:"id"`
	Username       string `json:"username"`
	DisplayName    string `json:"display_name"`
	CreatedAt      string `json:"created_at"`
	FollowersCount int    `json:"followers_count"`
	FollowingCount int    `json:"following_count"`
	Note           string `json:"note"`
	Avatar         string `json:"avatar"`
	Header         string `json:"header"`
}

// toPostUsersResponse: ユーザー新規登録のレスポンスに変換
func toPostUsersResponse(user *user.User) *PostUsersResponse {
	return &PostUsersResponse{
		ID:             user.ID(),
		Username:       user.Username(),
		DisplayName:    "",
		CreatedAt:      user.CreatedAt().Format("2006-01-02T15:04:05.000Z"),
		FollowersCount: 0,
		FollowingCount: 0,
		Note:           "",
		Avatar:         "",
		Header:         "",
	}
}
