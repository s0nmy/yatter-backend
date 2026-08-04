package user

import (
	"regexp"
	"time"
	"unicode/utf8"
	"yatter-backend-go/pkg/errors"

	"golang.org/x/crypto/bcrypt"
)

// Q. フィールド非公開にしてgetter/setterを作っているのはなぜ？
// A. フィールドの直接の操作を避けるため
//    NewUser, SetID などでのみ変更を行うことで常に整合成が保証されたインスタンスにしたい

// Q. フィールドがポインタでないのはなぜ？
// A. nilが入ることを避けるため
//    ポインタにするとnilが入る可能性を考慮する必要があり、意図しないパニックの原因にもなるためフィールドのポインタ化は避ける
//    ただ、オプショナルなフィールドや、別の構造体をフィールドに持ちたい場合にはポインタを使う
//        -> パフォーマンス面や、暗黙的な挙動の差異などでメリットがあるため
//        例. ディープコピーされてると思ったけど、実はフィールドのスライスは同じメモリを参照している、など

type User struct {
	id           uint64
	username     string
	passwordHash string
	createdAt    time.Time
}

// ReconstructUser で User を再構築する
// RepositoryでDBからロードする場合にこちらのファクトリメソッドを使うこと
func ReconstructUser(id uint64, username, passwordHash string, createdAt time.Time) (*User, error) {
	user := &User{}

	if err := user.SetID(id); err != nil {
		return nil, err
	}

	if err := user.SetUsername(username); err != nil {
		return nil, err
	}

	if err := user.SetPasswordHash(passwordHash); err != nil {
		return nil, err
	}

	if err := user.SetCreatedAt(createdAt); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) SetID(id uint64) error {
	// idは1以上の整数であること
	if id < 1 {
		return errors.ErrInternal.WithDevMessage("id must be more than 0")
	}

	u.id = id
	return nil
}

func (u *User) SetUsername(username string) error {
	// ユーザー名は1文字以上であること
	if utf8.RuneCountInString(username) < 1 {
		return errors.ErrInternal.WithDevMessage("user name must be more than 1 character")
	}

	// ユーザー名は50文字以下であること
	if utf8.RuneCountInString(username) > 50 {
		return errors.ErrInternal.WithDevMessage("user name must be less than 50 characters")
	}

	// ユーザー名は半角英数字とアンダースコアのみで構成されていること
	usernameRegexp := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	if !usernameRegexp.MatchString(username) {
		return errors.ErrInternal.WithDevMessage("user name must be alphanumeric and underscore")
	}

	u.username = username
	return nil
}

func (u *User) SetPasswordHash(password string) error {
	// パスワードは1文字以上であること
	if utf8.RuneCountInString(password) < 1 {
		return errors.ErrInternal.WithDevMessage("password must be more than 1 character")
	}

	// bcryptの仕様上、パスワードは 72byte 以下である必要がある
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.ErrInternal.WithDevMessage("failed to generate password hash")
	}

	u.passwordHash = string(hash)
	return nil
}

func (u *User) SetCreatedAt(createdAt time.Time) error {
	// createdAtはYatterサービス開始時以降であること
	yatterLaunchedAt := time.Date(2025, 4, 1, 0, 0, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60))
	if !createdAt.After(yatterLaunchedAt) {
		return errors.ErrInternal.WithDevMessage("createdAt must be after yatter launched")
	}

	// createdAtは未来の日付でないこと
	if createdAt.After(time.Now()) {
		return errors.ErrInternal.WithDevMessage("createdAt must not be in the future")
	}

	// createdAtのタイムゾーンはJSTであること
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	jstCreatedAt := createdAt.In(jst)
	if !createdAt.Equal(jstCreatedAt) {
		return errors.ErrInternal.WithDevMessage("createdAt must be in JST")
	}

	u.createdAt = createdAt

	return nil
}

func (u *User) ID() uint64 {
	return u.id
}

func (u *User) Username() string {
	return u.username
}

func (u *User) PasswordHash() string {
	return u.passwordHash
}

func (u *User) CreatedAt() time.Time {
	return u.createdAt
}
