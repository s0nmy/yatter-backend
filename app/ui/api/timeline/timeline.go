package timeline

import (
	"encoding/json"
	"fmt"
	"net/http"
	ui_errors "yatter-backend-go/app/ui/api/pkg/errors"
	timelineUseCase "yatter-backend-go/app/usecase/timeline"
	"yatter-backend-go/pkg/errors"
)

// ハンドラーのインターフェースを定義
type Handler interface {
	GetPublicTimeline(w http.ResponseWriter, r *http.Request)
}

func NewTimelineHandler(timelineGetPublicTimelineUseCase timelineUseCase.GetPublicTimelineUseCase) Handler {
	return &timelineHandlerImpl{
		timelineGetPublicTimelineUseCase: timelineGetPublicTimelineUseCase,
	}
}

var _ Handler = (*timelineHandlerImpl)(nil)

type timelineHandlerImpl struct {
	// ucを呼び出す
	timelineGetPublicTimelineUseCase timelineUseCase.GetPublicTimelineUseCase
}

func (h *timelineHandlerImpl) GetPublicTimeline(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// タイムラインusecase を呼び出し
	yweets, err := h.timelineGetPublicTimelineUseCase.GetPublicTimeline(ctx)
	if err != nil {
		ui_errors.Handle(w, err)
		return
	}

	responses := make([]*PublicTimelineYweetResponse, 0, len(yweets))
	for _, yweet := range yweets {
		responses = append(responses, toPublicTimelineYweetResponse(yweet))
	}

	// レスポンスのエンコードを行う
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(responses); err != nil {
		ui_errors.Handle(w, errors.ErrInternal.WithDevMessage(fmt.Sprintf("failed to encode response: %s", err.Error())))
		return
	}
}
