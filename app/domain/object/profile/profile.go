package profile

import (
	"net/url"
	"unicode/utf8"
	"yatter-backend-go/pkg/errors"
)

// ユーザーのプロフィール情報
type Profile struct {
	userID         uint64
	displayName    string
	avatarImageURL string
	headerImageURL string
	note           string
}

func NewProfile(userID uint64, displayName, avatarImageURL, headerImageURL, note string) (*Profile, error) {
	profile := &Profile{}

	if err := profile.SetUserID(userID); err != nil {
		return nil, err
	}

	if err := profile.SetDisplayName(displayName); err != nil {
		return nil, err
	}

	if err := profile.SetAvatarImageURL(avatarImageURL); err != nil {
		return nil, err
	}

	if err := profile.SetHeaderImageURL(headerImageURL); err != nil {
		return nil, err
	}

	if err := profile.SetNote(note); err != nil {
		return nil, err
	}

	return profile, nil
}

func (p *Profile) SetUserID(userID uint64) error {
	// userIDは1以上の整数であること
	if userID < 1 {
		return errors.ErrInternal.WithDevMessage("userID must be more than 0")
	}

	p.userID = userID
	return nil
}

func (p *Profile) SetDisplayName(displayName string) error {
	// 表示名は1文字以上であること
	if utf8.RuneCountInString(displayName) < 1 {
		return errors.ErrInternal.WithDevMessage("display name must be more than 0 characters")
	}

	// 表示名は50文字以下であること
	if utf8.RuneCountInString(displayName) > 50 {
		return errors.ErrInternal.WithDevMessage("display name must be less than or equal to 50 characters")
	}

	p.displayName = displayName
	return nil
}

func (p *Profile) SetAvatarImageURL(avatarImageURL string) error {
	// 正しいURL形式であること
	_, err := url.ParseRequestURI(avatarImageURL)
	if err != nil {
		return errors.ErrInternal.WithDevMessage("avatar image URL must be a valid URI")
	}

	// 今回は行わないが、セキュリティ上の観点から画像URLのホスト名が正しいかどうかもチェックする
	// 画像URLのホスト名はimage.yatter.comであること
	// if !(parsedURL.Hostname() == "image.yatter.com") {
	// 	return nil, fmt.Errorf("画像URLのホスト名が正しくありません")
	// }

	p.avatarImageURL = avatarImageURL
	return nil
}

func (p *Profile) SetHeaderImageURL(headerImageURL string) error {
	// 正しいURL形式であること
	_, err := url.ParseRequestURI(headerImageURL)
	if err != nil {
		return errors.ErrInternal.WithDevMessage("header image URL must be a valid URI")
	}

	// 今回は行わないが、セキュリティ上の観点から画像URLのホスト名が正しいかどうかもチェックする
	// 画像URLのホスト名はimage.yatter.comであること
	// if !(parsedURL.Hostname() == "image.yatter.com") {
	// 	return nil, fmt.Errorf("画像URLのホスト名が正しくありません")
	// }

	p.headerImageURL = headerImageURL
	return nil
}

func (p *Profile) SetNote(note string) error {
	// ノートは500文字以内であること
	if utf8.RuneCountInString(note) > 500 {
		return errors.ErrInternal.WithDevMessage("note must be less than or equal to 500 characters")
	}

	p.note = note
	return nil
}

func (p *Profile) UserID() uint64 {
	return p.userID
}

func (p *Profile) DisplayName() string {
	return p.displayName
}

func (p *Profile) AvatarImageURL() string {
	return p.avatarImageURL
}

func (p *Profile) HeaderImageURL() string {
	return p.headerImageURL
}

func (p *Profile) Note() string {
	return p.note
}
