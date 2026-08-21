package timeline

import (
	"yatter-backend-go/app/domain/object/yweet"
)

// 公開タイムラインに含まれる Yweet 1件のレスポンスである。
type PublicTimelineYweetResponse struct {
	ID uint64 `json:"id"`
	// User      *user.PostUsersResponse `json:"user"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

func toPublicTimelineYweetResponse(yweet *yweet.Yweet) *PublicTimelineYweetResponse {
	return &PublicTimelineYweetResponse{
		ID: yweet.ID(),
		// User:      user,
		Content:   yweet.Content(),
		CreatedAt: yweet.CreatedAt().Format("2006-01-02T15:04:05.000Z"),
	}
}
