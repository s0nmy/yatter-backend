package yweet

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	userObject "yatter-backend-go/app/domain/object/user"
	ui_errors "yatter-backend-go/app/ui/api/pkg/errors"
	handlerUser "yatter-backend-go/app/ui/api/user"
	yweetUseCase "yatter-backend-go/app/usecase/yweet"
	"yatter-backend-go/pkg/errors"
)

// テストしやすいように、ハンドラーのインターフェースを定義
type Handler interface {
	CreateYweet(w http.ResponseWriter, r *http.Request)
}

func NewYweetHandler(yweetCreateUseCase yweetUseCase.CreateYweetUseCase) Handler {
	return &yweetHandlerImpl{
		// 1つ下のucを持ってくる
		yweetCreateUseCase: yweetCreateUseCase,
	}
}

var _ Handler = (*yweetHandlerImpl)(nil)

// yweetHandler はユーザー関連の API を管理
type yweetHandlerImpl struct {
	yweetCreateUseCase yweetUseCase.CreateYweetUseCase
}

// yweetを
func (h *yweetHandlerImpl) CreateYweet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// リクエストをデコード
	// username の取得
	fmt.Println(r.Header.Get("Authentication"))

	// username を受け取って格納する
	username := strings.Split(r.Header.Get("Authentication"), " ")[1]

	var req PostYweetRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ui_errors.Handle(w, errors.ErrBadRequest)
		return
	}

	// yweet投稿ユースケースを呼び出し
	createdYweet, createdUser, err := h.yweetCreateUseCase.CreateYweet(ctx, username, req.Content)
	if err != nil {
		ui_errors.Handle(w, err)
		return
	}

	// レスポンスに変換
	resp := toPostYweetResponse(createdYweet, toPostUserResponse(createdUser))

	// レスポンスをエンコードして返す
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(resp); err != nil {
		ui_errors.Handle(w, errors.ErrInternal.WithDevMessage(fmt.Sprintf("failed to encode response: %s", err.Error())))
		return
	}
}

func toPostUserResponse(usr *userObject.User) *handlerUser.PostUsersResponse {
	return &handlerUser.PostUsersResponse{
		ID:             usr.ID(),
		Username:       usr.Username(),
		DisplayName:    "",
		CreatedAt:      usr.CreatedAt().Format("2006-01-02T15:04:05.000Z"),
		FollowersCount: 0,
		FollowingCount: 0,
		Note:           "",
		Avatar:         "",
		Header:         "",
	}
}
