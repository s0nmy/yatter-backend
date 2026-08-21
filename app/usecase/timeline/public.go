package timeline

import (
	"context"
	"errors"
	"yatter-backend-go/app/domain/object/yweet"
	"yatter-backend-go/app/domain/repository"
	"yatter-backend-go/app/usecase/transactor"
)

type GetPublicTimelineUseCase interface {
	GetPublicTimeline(ctx context.Context) ([]*yweet.Yweet, error)
}

var _ GetPublicTimelineUseCase = (*getPublicTimelineUseCaseImpl)(nil)

type getPublicTimelineUseCaseImpl struct {
	yweetRepo  repository.Yweet
	transactor transactor.Transactor
}

func NewGetPublicTimelineUseCase(
	yweetRepo repository.Yweet,
	transactor transactor.Transactor,
) *getPublicTimelineUseCaseImpl {
	return &getPublicTimelineUseCaseImpl{
		yweetRepo:  yweetRepo,
		transactor: transactor,
	}
}

func (uc *getPublicTimelineUseCaseImpl) GetPublicTimeline(ctx context.Context) ([]*yweet.Yweet, error) {
	// トランザクションを張る
	result, err := uc.transactor.TransactionWithValue(ctx, func(ctx context.Context) (any, error) {
		// yweetを全件取得
		return uc.yweetRepo.SelectAll(ctx)
	})
	if err != nil {
		return nil, err
	}

	// 複数件を格納するため、スライスを利用することを忘れない
	publicTimeline, ok := result.([]*yweet.Yweet)
	if !ok {
		return nil, errors.New("failed to cast result to []*yweet.Yweet")
	}

	return publicTimeline, nil
}
