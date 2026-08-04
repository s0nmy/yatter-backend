package user_test

import (
	"strings"
	"testing"
	"yatter-backend-go/app/domain/object/user"
	yatter_errors "yatter-backend-go/pkg/errors"
)

func Test_PendingUser_SetUsername(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		username string
		wantErr  error
	}{
		{
			name:     "正常系: ユーザー名が1文字の半角英数字の場合、ユーザー名がセットされ何も返されない",
			username: "a",
			wantErr:  nil,
		},
		{
			name:     "正常系: ユーザー名が50文字の半角英数字とアンダースコアの場合、ユーザー名がセットされ何も返されない",
			username: strings.Repeat("a", 50),
			wantErr:  nil,
		},
		{
			name:     "正常系: ユーザー名に大文字、小文字、数字、アンダースコアが含まれる場合、ユーザー名がセットされ何も返されない",
			username: "Test_User123",
			wantErr:  nil,
		},
		{
			name:     "異常系: ユーザー名が空文字列の場合、エラーが返される",
			username: "",
			wantErr:  yatter_errors.ErrBadRequest,
		},
		{
			name:     "異常系: ユーザー名が51文字の場合、エラーが返される",
			username: strings.Repeat("a", 51),
			wantErr:  yatter_errors.ErrBadRequest,
		},
		{
			name:     "異常系: ユーザー名にスペースが含まれる場合、エラーが返される",
			username: "test user",
			wantErr:  yatter_errors.ErrBadRequest,
		},
		{
			name:     "異常系: ユーザー名に日本語が含まれる場合、エラーが返される",
			username: "テスト",
			wantErr:  yatter_errors.ErrBadRequest,
		},
		{
			name:     "異常系: ユーザー名に特殊記号が含まれる場合、エラーが返される",
			username: "test-user",
			wantErr:  yatter_errors.ErrBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			u := &user.PendingUser{}

			err := u.SetUsername(tt.username)

			if !yatter_errors.Is(err, tt.wantErr) {
				t.Errorf("SetUsername() error = %v, wantErr %v", err, tt.wantErr)
			}

			// 正常系の場合は、ユーザー名が正しくセットされているか確認
			if tt.wantErr == nil {
				if u.Username() != tt.username {
					t.Errorf("Username() = %v, want %v", u.Username(), tt.username)
				}
			}
		})
	}
}

func Test_PendingUser_SetPasswordHash(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		password string
		wantErr  error
	}{
		{
			name:     "正常系: パスワードが1文字の場合、パスワードハッシュがセットされ何も返されない",
			password: "a",
			wantErr:  nil,
		},
		{
			name:     "正常系: パスワードが通常の長さの場合、パスワードハッシュがセットされ何も返されない",
			password: "password123",
			wantErr:  nil,
		},
		{
			name:     "正常系: パスワードに特殊文字が含まれる場合、パスワードハッシュがセットされ何も返されない",
			password: "P@ssw0rd!",
			wantErr:  nil,
		},
		{
			name:     "正常系: パスワードに日本語が含まれる場合、パスワードハッシュがセットされ何も返されない",
			password: "パスワード123",
			wantErr:  nil,
		},
		{
			name:     "異常系: パスワードが空文字列の場合、エラーが返される",
			password: "",
			wantErr:  yatter_errors.ErrBadRequest,
		},
		{
			name:     "正常系: 長いパスワード（72バイト以内）の場合、パスワードハッシュがセットされ何も返されない",
			password: strings.Repeat("a", 72),
			wantErr:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			u := &user.PendingUser{}

			err := u.SetPasswordHash(tt.password)

			if !yatter_errors.Is(err, tt.wantErr) {
				t.Errorf("SetPasswordHash() error = %v, wantErr %v", err, tt.wantErr)
			}

			// 正常系の場合は、パスワードハッシュが正しくセットされているか確認
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
