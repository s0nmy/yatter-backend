package repository

import (
	"context"
	"yatter-backend-go/app/domain/object/user"
)

// aggregate 単位で更新を行う
// Find系の引数には値オブジェクトなどを使ってもよい
type User interface {
	// 指定したユーザー名のユーザーを取得する
	// ユーザーが存在しない場合は nil, nil を返す
	FindByUsername(ctx context.Context, username string) (*user.User, error)

	// ユーザーを新規作成して保存する
	// 保存に成功した場合は保存したユーザーを返す
	Insert(ctx context.Context, pendingUser *user.PendingUser) (*user.User, error)
}
