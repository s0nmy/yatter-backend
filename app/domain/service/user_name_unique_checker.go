package service

import (
	"context"
	"yatter-backend-go/app/domain/repository"
)

type UsernameUniqueChecker struct {
	userRepository repository.User
}

func NewUsernameUniqueChecker(userRepository repository.User) *UsernameUniqueChecker {
	return &UsernameUniqueChecker{
		userRepository: userRepository,
	}
}

func (c *UsernameUniqueChecker) IsUnique(ctx context.Context, username string) (bool, error) {
	user, err := c.userRepository.FindByUsername(ctx, username)
	if err != nil {
		return false, err
	}

	// ユーザーが存在しない場合はユーザー名が重複していない
	return user == nil, nil
}
