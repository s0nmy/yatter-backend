package auth

import (
	"regexp"
	"unicode/utf8"
	"yatter-backend-go/pkg/errors"
)

// Credential は認証情報を表す構造体
// ユーザー名とパスワードハッシュを持つ
// DBから読み取るときなどはこの構造体を使う
type Credential struct {
	username     string
	passwordHash string
}

func ReconstructCredential(username, passwordHash string) (*Credential, error) {
	credential := &Credential{}

	if err := credential.SetUsername(username); err != nil {
		return nil, err
	}

	if err := credential.SetPasswordHash(passwordHash); err != nil {
		return nil, err
	}

	return credential, nil
}

func (c *Credential) SetUsername(username string) error {
	// ユーザー名は1文字以上であること
	if utf8.RuneCountInString(username) < 1 {
		return errors.ErrBadRequest
	}

	// ユーザー名は50文字以下であること
	if utf8.RuneCountInString(username) > 50 {
		return errors.ErrBadRequest
	}

	// ユーザー名は半角英数字とアンダースコアのみで構成されていること
	usernameRegexp := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	if !usernameRegexp.MatchString(username) {
		return errors.ErrBadRequest
	}

	c.username = username
	return nil
}

func (c *Credential) SetPasswordHash(passwordHash string) error {
	// パスワードは空でないこと
	if len(passwordHash) == 0 {
		return errors.ErrBadRequest
	}

	c.passwordHash = passwordHash
	return nil
}

func (c *Credential) Username() string {
	return c.username
}

func (c *Credential) PasswordHash() string {
	return c.passwordHash
}
