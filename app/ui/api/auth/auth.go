package auth

import (
	"encoding/json"
	"net/http"
	ui_errors "yatter-backend-go/app/ui/api/pkg/errors"
	"yatter-backend-go/app/usecase/auth"
	"yatter-backend-go/pkg/errors"
)

// テストしやすいように、ハンドラーのインターフェースを定義
type Handler interface {
	Login(w http.ResponseWriter, r *http.Request)
}

func NewAuthHandler(auth auth.LoginUseCase) Handler {
	return &authHandlerImpl{
		auth: auth,
	}
}

var _ Handler = (*authHandlerImpl)(nil)

// authHandler は認証関連の API を管理
type authHandlerImpl struct {
	auth auth.LoginUseCase
}

// Login: ユーザーログイン
func (h *authHandlerImpl) Login(w http.ResponseWriter, r *http.Request) {
	// リクエストをデコード
	var req PostLoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ui_errors.Handle(w, errors.ErrBadRequest)
		return
	}

	ctx := r.Context()

	// ログインユースケースを呼び出し
	credential, err := h.auth.Login(ctx, req.Username, req.Password)
	if err != nil {
		ui_errors.Handle(w, err)
		return
	}

	// レスポンスに変換
	resp := toPostLoginResponse(credential)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(resp); err != nil {
		ui_errors.Handle(w, errors.ErrInternal.WithDevMessage("failed to encode response"))
		return
	}
}
