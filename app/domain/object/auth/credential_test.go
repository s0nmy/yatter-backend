package auth_test

import (
	"strings"
	"testing"
	"yatter-backend-go/app/domain/object/auth"
	yatter_errors "yatter-backend-go/pkg/errors"
)

func Test_Credential_SetUsername(t *testing.T) {
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

			c := &auth.Credential{}

			err := c.SetUsername(tt.username)

			if !yatter_errors.Is(err, tt.wantErr) {
				t.Errorf("SetUsername() error = %v, wantErr %v", err, tt.wantErr)
			}

			// 正常系の場合は、ユーザー名が正しくセットされているか確認
			if tt.wantErr == nil {
				if c.Username() != tt.username {
					t.Errorf("Username() = %v, want %v", c.Username(), tt.username)
				}
			}
		})
	}
}

func Test_Credential_SetPasswordHash(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		passwordHash string
		wantErr      error
	}{
		{
			name:         "正常系: パスワードハッシュが通常の文字列の場合、パスワードハッシュがセットされ何も返されない",
			passwordHash: "hashed_password123",
			wantErr:      nil,
		},
		{
			name:         "正常系: パスワードハッシュが長い文字列の場合、パスワードハッシュがセットされ何も返されない",
			passwordHash: strings.Repeat("a", 100),
			wantErr:      nil,
		},
		{
			name:         "異常系: パスワードハッシュが空文字列の場合、エラーが返される",
			passwordHash: "",
			wantErr:      yatter_errors.ErrBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			c := &auth.Credential{}

			err := c.SetPasswordHash(tt.passwordHash)

			if !yatter_errors.Is(err, tt.wantErr) {
				t.Errorf("SetPasswordHash() error = %v, wantErr %v", err, tt.wantErr)
			}

			// 正常系の場合は、パスワードハッシュが正しくセットされているか確認
			if tt.wantErr == nil {
				if c.PasswordHash() != tt.passwordHash {
					t.Errorf("PasswordHash() = %v, want %v", c.PasswordHash(), tt.passwordHash)
				}
			}
		})
	}
}
