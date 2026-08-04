package user

import (
	"context"
	"yatter-backend-go/app/domain/object/user"
	"yatter-backend-go/app/domain/repository"
	"yatter-backend-go/app/domain/service"
	"yatter-backend-go/app/usecase/transactor"
	"yatter-backend-go/pkg/errors"
)

type SignUpUseCase interface {
	SignUp(ctx context.Context, username, password string) (*user.User, error)
}

var _ SignUpUseCase = (*SignUpUseCaseImpl)(nil)

type SignUpUseCaseImpl struct {
	userRepo              repository.User
	usernameUniqueChecker *service.UsernameUniqueChecker
	transactor            transactor.Transactor
}

func NewUserCreateUseCase(
	userRepo repository.User,
	usernameUniqueChecker *service.UsernameUniqueChecker,
	transactor transactor.Transactor,
) *SignUpUseCaseImpl {
	return &SignUpUseCaseImpl{
		userRepo:              userRepo,
		usernameUniqueChecker: usernameUniqueChecker,
		transactor:            transactor,
	}
}

func (uc *SignUpUseCaseImpl) SignUp(ctx context.Context, username, password string) (*user.User, error) {
	result, err := uc.transactor.TransactionWithValue(ctx, func(ctx context.Context) (any, error) {
		// ユーザー名が重複しているかを取得
		isUniqueUsername, err := uc.usernameUniqueChecker.IsUnique(ctx, username)
		if err != nil {
			return nil, err
		}

		// 仮登録ユーザーを生成
		pendingUser, err := user.NewPendingUser(username, password, isUniqueUsername)
		if err != nil {
			return nil, err
		}

		// 仮登録ユーザーを保存
		user, err := uc.userRepo.Insert(ctx, pendingUser)
		if err != nil {
			return nil, err
		}

		return user, nil
	})
	if err != nil {
		return nil, err
	}

	user, ok := result.(*user.User)
	if !ok {
		return nil, errors.ErrInternal.WithDevMessage("failed to cast result to user.User")
	}

	return user, nil
}
