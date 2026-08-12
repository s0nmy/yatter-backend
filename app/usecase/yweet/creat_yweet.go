package yweetUseCase

import (
	"context"
	"time"
	"yatter-backend-go/app/domain/object/user"
	"yatter-backend-go/app/domain/object/yweet"

	"yatter-backend-go/app/domain/repository"
	"yatter-backend-go/app/usecase/transactor"
	"yatter-backend-go/pkg/errors"
)

// 概念を呼びたいので、具体的なもの(構造体)を作りたくない

type CreateYweetUseCase interface {
	CreateYweet(ctx context.Context, username string, content string) (*yweet.Yweet, error)
}

var _ CreateYweetUseCase = (*createYweetUseCaseImpl)(nil)

type createYweetUseCaseImpl struct {
	userRepo   repository.User
	yweetRepo  repository.Yweet
	transactor transactor.Transactor
}

func NewYweetCreateUseCase(
	userRepo repository.User,
	yweetRepo repository.Yweet,
	transactor transactor.Transactor,
) *createYweetUseCaseImpl {
	return &createYweetUseCaseImpl{
		userRepo:   userRepo,
		yweetRepo:  yweetRepo,
		transactor: transactor,
	}
}

func (uc *createYweetUseCaseImpl) CreateYweet(ctx context.Context, username string, content string) (*yweet.Yweet, *user.User, error) {
	// トランザクション処理
	result, err := uc.transactor.TransactionWithValue(ctx, func(ctx context.Context) (any, error) {

		userID, err := uc.userRepo.FindByUsername(ctx, username)

		yweet, err := yweet.NewYweet(
			0,
			userID.ID(),
			content,
			time.Time{})
		if err != nil {
			return nil, err
		}

		InsertedYweet, err := uc.yweetRepo.Insert(ctx, yweet)
		if err != nil {
			return nil, err
		}

		user

		return InsertedYweet, nil
	})
	if err != nil {
		return nil, nil, err
	}

	yweet, ok := result.(*yweet.Yweet)
	if !ok {
		return nil, nil, errors.ErrInternal.WithDevMessage("failed to cast result to yweet.Yweet")
	}

	return yweet, user, nil
}
