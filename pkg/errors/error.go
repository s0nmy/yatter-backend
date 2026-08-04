package errors

import (
	"errors"
	"yatter-backend-go/pkg/errors/code"
)

var (
	ErrBadRequest   = New(code.BadRequest, "bad request", "")                                       // NOTE: 不正なリクエスト（例: リクエストがJSON形式になっていない）
	ErrUnauthorized = New(code.Unauthorized, "unauthorized", "")                                    // NOTE: 認証エラー（例: Authorizationヘッダーが設定されていない）
	ErrForbidden    = New(code.Forbidden, "forbidden", "")                                          // NOTE: 認可エラー（例: 認証したユーザーと異なるユーザーの投稿を削除しようとしている）
	ErrNotFound     = New(code.NotFound, "not found", "")                                           // NOTE: リソースが見つからない（例: 指定されたIDの投稿が見つからない）
	ErrConflict     = New(code.Conflict, "conflict", "")                                            // NOTE: リソースの重複（例: ユーザー名がすでに利用されている）
	ErrInternal     = New(code.Internal, "internal server error", "internal server error occurred") // NOTE: サーバーエラー（例: データベースの接続エラー）
)

/*
研修資料
-----

詳細なエラーは作成せず作成済みの汎用エラーを利用しても良いが、余力がある場合は以下のような詳細なエラーを作成する.
ErrUsernameTooLong = New(code.BadRequest, "bad request", "user name must be less than 50 characters")

このファイルに作成しても良いし、/app/domain/object/user/user.go などリソースのファイルに作成しても良い.
*/

func Is(err1, err2 error) bool {
	status1, ok := err1.(*Status)
	if !ok {
		return errors.Is(err1, err2)
	}

	status2, ok := err2.(*Status)
	if !ok {
		return errors.Is(err1, err2)
	}

	return status1.Code() == status2.Code()
}
