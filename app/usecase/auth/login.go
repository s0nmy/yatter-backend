package auth

import (
	"context"
	"yatter-backend-go/app/domain/object/auth"
	"yatter-backend-go/app/usecase/query"
)

type LoginUseCase interface {
	Login(ctx context.Context, username, password string) (*auth.AuthenticatedCredential, error)
}

var _ LoginUseCase = (*LoginUseCaseImpl)(nil)

type LoginUseCaseImpl struct {
	authQueryService query.AuthQueryService
}

func NewLoginUseCase(authQueryService query.AuthQueryService) *LoginUseCaseImpl {
	return &LoginUseCaseImpl{
		authQueryService: authQueryService,
	}
}

func (uc *LoginUseCaseImpl) Login(ctx context.Context, username, password string) (*auth.AuthenticatedCredential, error) {
	credential, err := uc.authQueryService.FindCredentialByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	authenticatedCredential, err := auth.NewAuthenticatedCredential(credential, password)
	if err != nil {
		return nil, err
	}

	return authenticatedCredential, nil
}
