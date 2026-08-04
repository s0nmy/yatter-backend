package user_test

import (
	"strings"
	"testing"
	"yatter-backend-go/app/domain/object/user"
	yatter_errors "yatter-backend-go/pkg/errors"
)

func Test_User_SetID(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		id      uint64
		wantErr error
	}{
		{
			name:    "正常系: idが1以上の整数の場合、IDがセットされ何も返されない",
			id:      1,
			wantErr: nil,
		},
		{
			name:    "正常系: idが2^64の場合、IDがセットされ何も返されない",
			id:      ^uint64(0),
			wantErr: nil,
		},
		{
			name:    "異常系: idが0の場合、エラーが返される",
			id:      0,
			wantErr: yatter_errors.ErrInternal,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			u := &user.User{}

			err := u.SetID(tt.id)

			if !yatter_errors.Is(err, tt.wantErr) {
				t.Errorf("SetID() error = %v, wantErr %v", err, tt.wantErr)
			}

			// 正常系の場合は、IDがセットされていることを確認
			if tt.wantErr == nil {
				if u.ID() != tt.id {
					t.Errorf("ID() = %v, want %v", u.ID(), tt.id)
				}
			}
		})
	}
}

func Test_User_SetUsername(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		username string
		wantErr  error
	}{
		{
			name:     "正常系: ユーザー名が1文字以上50文字以下の半角英数字またはアンダースコアである場合、ユーザー名がセットされ何も返されない",
			username: "test_user",
			wantErr:  nil,
		},
		{
			name:     "異常系: ユーザー名が0文字の場合、エラーが返される",
			username: "",
			wantErr:  yatter_errors.ErrInternal,
		},
		{
			name:     "異常系: ユーザー名が50文字を超える場合、エラーが返される",
			username: strings.Repeat("a", 51),
			wantErr:  yatter_errors.ErrInternal,
		},
		{
			name:     "異常系: ユーザー名が半角英数字またはアンダースコア以外の文字を含む場合、エラーが返される",
			username: "test user",
			wantErr:  yatter_errors.ErrInternal,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			u := &user.User{}

			err := u.SetUsername(tt.username)

			if !yatter_errors.Is(err, tt.wantErr) {
				t.Errorf("SetName() error = %v, wantErr %v", err, tt.wantErr)
			}

			// 正常系の場合は、ユーザー名がセットされていることを確認
			if tt.wantErr == nil {
				if u.Username() != tt.username {
					t.Errorf("Username() = %v, want %v", u.Username(), tt.username)
				}
			}

		})
	}
}

func Test_User_SetPasswordHash(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		password string
		wantErr  error
	}{
		{
			name:     "正常系: パスワードが1文字以上かつ72byte以下の場合、パスワードハッシュがセットされ何も返されない",
			password: "password123",
			wantErr:  nil,
		},
		{
			name:     "異常系: パスワードが空文字列の場合、エラーが返される",
			password: "",
			wantErr:  yatter_errors.ErrInternal,
		},
		{
			name:     "異常系: パスワードが72byteを超える場合、エラーが返される",
			password: "a" + string(make([]byte, 72)), // 73 bytes
			wantErr:  yatter_errors.ErrInternal,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			u := &user.User{}

			err := u.SetPasswordHash(tt.password)

			if !yatter_errors.Is(err, tt.wantErr) {
				t.Errorf("SetPasswordHash() error = %v, wantErr %v", err, tt.wantErr)
			}

			// 正常系の場合は、パスワードハッシュがセットされていることを確認
			if tt.wantErr == nil {
				if u.PasswordHash() == "" {
					t.Error("PasswordHash() is empty, but it should be set")
				}

				// パスワードとハッシュ値が異なることを確認（ハッシュ化されていることを確認）
				if u.PasswordHash() == tt.password {
					t.Error("PasswordHash() should be different from the original password")
				}
			}
		})
	}
}
