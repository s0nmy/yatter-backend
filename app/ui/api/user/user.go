package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	ui_errors "yatter-backend-go/app/ui/api/pkg/errors"
	"yatter-backend-go/app/usecase/user"
	"yatter-backend-go/pkg/errors"
)

// テストしやすいように、ハンドラーのインターフェースを定義
type Handler interface {
	SignUp(w http.ResponseWriter, r *http.Request)
}

func NewUserHandler(userCreateUseCase user.SignUpUseCase) Handler {
	return &userHandlerImpl{
		userCreateUseCase: userCreateUseCase,
	}
}

var _ Handler = (*userHandlerImpl)(nil)

// userHandler はユーザー関連の API を管理
type userHandlerImpl struct {
	userCreateUseCase user.SignUpUseCase
}

// SignUpUser: ユーザー新規登録
func (h *userHandlerImpl) SignUp(w http.ResponseWriter, r *http.Request) {
	// リクエストをデコード
	var req PostUsersRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ui_errors.Handle(w, errors.ErrBadRequest)
		return
	}

	ctx := r.Context()

	// 新規登録ユースケースを呼び出し
	signedUpUser, err := h.userCreateUseCase.SignUp(ctx, req.Username, req.Password)
	if err != nil {
		ui_errors.Handle(w, err)
		return
	}

	// レスポンスに変換
	resp := toPostUsersResponse(signedUpUser)

	// レスポンスをエンコードして返す
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(resp); err != nil {
		ui_errors.Handle(w, errors.ErrInternal.WithDevMessage(fmt.Sprintf("failed to encode response: %s", err.Error())))
		return
	}
}
