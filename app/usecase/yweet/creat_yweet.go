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
	CreateYweet(ctx context.Context, username string, content string) (*yweet.Yweet, *user.User, error)
}

var _ CreateYweetUseCase = (*createYweetUseCaseImpl)(nil)

type createYweetUseCaseImpl struct {
	userRepo   repository.User
	yweetRepo  repository.Yweet
	transactor transactor.Transactor
}

// 投稿自体を1つの構造体とする
type createYweetResult struct {
	yweet *yweet.Yweet
	user  *user.User
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

		// usr は *user.User を受け取っている
		usr, err := uc.userRepo.FindByUsername(ctx, username)
		if err != nil {
			return nil, err
		}

		newYweet, err := yweet.NewYweet(
			0,
			usr.ID(),
			content,
			time.Time{})
		if err != nil {
			return nil, err
		}

		insertedYweet, err := uc.yweetRepo.Insert(ctx, newYweet)
		if err != nil {
			return nil, err
		}

		return &createYweetResult{yweet: insertedYweet, user: usr}, nil
	})
	if err != nil {
		return nil, nil, err
	}

	created, ok := result.(*createYweetResult)
	if !ok {
		return nil, nil, errors.ErrInternal.WithDevMessage("failed to cast result to createYweetResult")
	}

	return created.yweet, created.user, nil
}
