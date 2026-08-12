package yweet

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	ui_errors "yatter-backend-go/app/ui/api/pkg/errors"
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
	createdYweet, err := h.yweetCreateUseCase.CreateYweet(ctx, username, req.Content)
	if err != nil {
		ui_errors.Handle(w, err)
		return
	}

	// レスポンスに変換
	resp := toPostYweetResponse(createdYweet, user)

	// レスポンスをエンコードして返す
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(resp); err != nil {
		ui_errors.Handle(w, errors.ErrInternal.WithDevMessage(fmt.Sprintf("failed to encode response: %s", err.Error())))
		return
	}
}
