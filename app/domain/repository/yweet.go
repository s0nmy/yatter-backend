package repository

import (
	"context"
	"yatter-backend-go/app/domain/object/yweet"
)

// aggregate 単位で更新を行う
// Find系の引数には値オブジェクトなどを使ってもよい
type Yweet interface {
	// ユーザーを新規作成して保存する
	// 保存に成功した場合は保存したユーザーを返す
	Insert(ctx context.Context, Yweet *yweet.Yweet) (*yweet.Yweet, error)
	// timeline の取得にあたり、全件取得用を返す
	SelectAll(ctx context.Context) ([]*yweet.Yweet, error)
}
