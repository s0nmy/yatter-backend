package user

import (
	"regexp"
	"unicode/utf8"
	"yatter-backend-go/pkg/errors"

	"golang.org/x/crypto/bcrypt"
)

type PendingUser struct {
	username     string
	passwordHash string
}

// NewPendingUser で 仮登録ユーザー を新規作成する
// ユーザーを新規登録する場合こちらのファクトリメソッドを使うこと
//
// NOTE:
// bool値でユーザー名の重複を受け取っているのは、ユーザー新規生成時のロジックとして「ユーザー名が重複していたら新規登録できない」という挙動を実現するため
func NewPendingUser(username, password string, isUniqueUsername bool) (*PendingUser, error) {
	if !isUniqueUsername {
		return nil, errors.ErrConflict
	}

	user := &PendingUser{}

	if err := user.SetUsername(username); err != nil {
		return nil, err
	}

	if err := user.SetPasswordHash(password); err != nil {
		return nil, err
	}

	return user, nil
}

// 研修用にSetXxxのような形にしたが、同じ目的であるユーザー名のバリデーションコードが重複してしまっている
// （値オブジェクトは複雑 -> SetXxx という形で決まりましたがやっぱり値オブジェクトを使って実装するのが良いかもしれない）

func (u *PendingUser) SetUsername(username string) error {
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

	u.username = username
	return nil
}

func (u *PendingUser) SetPasswordHash(password string) error {
	// パスワードは1文字以上であること
	if utf8.RuneCountInString(password) < 1 {
		return errors.ErrBadRequest
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.ErrBadRequest
	}
	u.passwordHash = string(hash)
	return nil
}

func (u *PendingUser) Username() string {
	return u.username
}

func (u *PendingUser) PasswordHash() string {
	return u.passwordHash
}
