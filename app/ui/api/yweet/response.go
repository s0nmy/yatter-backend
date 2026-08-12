package yweet

import (
	"yatter-backend-go/app/domain/object/yweet"
	"yatter-backend-go/app/ui/api/user"
)

type PostYweetResponse struct {
	ID        uint64                  `json:"id"`
	User      *user.PostUsersResponse `json:"user"`
	Content   string                  `json:"content"`
	CreatedAt string                  `json:"created_at"`
}

// toPostYweetResponse: yweetの投稿用レスポンスに変換
func toPostYweetResponse(yweet *yweet.Yweet, user *user.PostUsersResponse) *PostYweetResponse {
	return &PostYweetResponse{
		ID:        yweet.ID(),
		User:      user,
		Content:   yweet.Content(),
		CreatedAt: yweet.CreatedAt().Format("2006-01-02T15:04:05.000Z"),
	}
}
