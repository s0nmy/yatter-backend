package auth_test

import (
	"testing"
	"yatter-backend-go/app/domain/object/auth"

	"golang.org/x/crypto/bcrypt"
)

func Test_NewAuthenticatedCredential(t *testing.T) {
	t.Parallel()

	// テスト用のパスワードとそのハッシュ値を作成
	password := "password123"
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("Failed to generate password hash: %v", err)
	}

	tests := []struct {
		name         string
		password     string
		wantUsername string
		wantErr      bool
	}{
		{
			name:         "正常系: 正しいパスワードの場合、AuthenticatedCredentialが生成される",
			password:     password,
			wantUsername: "testuser",
			wantErr:      false,
		},
		{
			name:         "異常系: 不正なパスワードの場合、エラーが返される",
			password:     "wrong_password",
			wantUsername: "",
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			credential, err := auth.ReconstructCredential("testuser", string(passwordHash))
			if err != nil {
				// テストケースの不備なので、テストを即時に失敗させる
				t.Fatalf("Failed to create credential: %v", err)
			}

			authenticatedCredential, err := auth.NewAuthenticatedCredential(credential, tt.password)

			// エラーの有無をチェック
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAuthenticatedCredential() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// 正常系の場合は、ユーザー名とパスワードハッシュが正しくセットされているか確認
			if err == nil {
				if authenticatedCredential.Username() != tt.wantUsername {
					t.Errorf("Username() = %v, want %v", authenticatedCredential.Username(), tt.wantUsername)
				}

				// パスワードハッシュが正しくセットされているか確認
				if authenticatedCredential.PasswordHash() != credential.PasswordHash() {
					t.Errorf("PasswordHash() = %v, want %v", authenticatedCredential.PasswordHash(), credential.PasswordHash())
				}
			}
		})
	}
}
