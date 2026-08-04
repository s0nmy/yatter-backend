package profile_test

import (
	"strings"
	"testing"
	"yatter-backend-go/app/domain/object/profile"
	yatter_errors "yatter-backend-go/pkg/errors"
)

func Test_Profile_SetUserID(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		userID  uint64
		wantErr error
	}{
		{
			name:    "正常系: ユーザーIDが1の場合、ユーザーIDがセットされ何も返されない",
			userID:  1,
			wantErr: nil,
		},
		{
			name:    "正常系: ユーザーIDが最大値の場合、ユーザーIDがセットされ何も返されない",
			userID:  ^uint64(0),
			wantErr: nil,
		},
		{
			name:    "異常系: ユーザーIDが0の場合、エラーが返される",
			userID:  0,
			wantErr: yatter_errors.ErrInternal,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			p := &profile.Profile{}

			err := p.SetUserID(tt.userID)

			if !yatter_errors.Is(err, tt.wantErr) {
				t.Errorf("SetUserID() error = %v, wantErr %v", err, tt.wantErr)
			}

			// 正常系の場合は、ユーザーIDが正しくセットされているか確認
			if tt.wantErr == nil {
				if p.UserID() != tt.userID {
					t.Errorf("UserID() = %v, want %v", p.UserID(), tt.userID)
				}
			}
		})
	}
}

func Test_Profile_SetDisplayName(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		displayName string
		wantErr     error
	}{
		{
			name:        "正常系: 表示名が1文字の場合、表示名がセットされ何も返されない",
			displayName: "A",
			wantErr:     nil,
		},
		{
			name:        "正常系: 表示名が50文字の場合、表示名がセットされ何も返されない",
			displayName: "あいうえおあいうえおあいうえおあいうえおあいうえおあいうえおあいうえおあいうえおあいうえおあいうえお",
			wantErr:     nil,
		},
		{
			name:        "正常系: 表示名に日本語が含まれる場合、表示名がセットされ何も返されない",
			displayName: "田中太郎",
			wantErr:     nil,
		},
		{
			name:        "正常系: 表示名に特殊文字が含まれる場合、表示名がセットされ何も返されない",
			displayName: "John Doe #1",
			wantErr:     nil,
		},
		{
			name:        "異常系: 表示名が空文字列の場合、エラーが返される",
			displayName: "",
			wantErr:     yatter_errors.ErrInternal,
		},
		{
			name:        "異常系: 表示名が51文字の場合、エラーが返される",
			displayName: "あいうえおあいうえおあいうえおあいうえおあいうえおあいうえおあいうえおあいうえおあいうえおあいうえおあ",
			wantErr:     yatter_errors.ErrInternal,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			p := &profile.Profile{}

			err := p.SetDisplayName(tt.displayName)

			if !yatter_errors.Is(err, tt.wantErr) {
				t.Errorf("SetDisplayName() error = %v, wantErr %v", err, tt.wantErr)
			}

			// 正常系の場合は、表示名が正しくセットされているか確認
			if tt.wantErr == nil {
				if p.DisplayName() != tt.displayName {
					t.Errorf("DisplayName() = %v, want %v", p.DisplayName(), tt.displayName)
				}
			}
		})
	}
}

func Test_Profile_SetAvatarImageURL(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		avatarImageURL string
		wantErr        error
	}{
		{
			name:           "正常系: 正しいURL形式の場合、アバター画像URLがセットされ何も返されない",
			avatarImageURL: "https://example.com/avatar.png",
			wantErr:        nil,
		},
		{
			name:           "正常系: HTTPスキームのURLの場合、アバター画像URLがセットされ何も返されない",
			avatarImageURL: "http://example.com/avatar.jpg",
			wantErr:        nil,
		},
		{
			name:           "正常系: クエリパラメータを含むURLの場合、アバター画像URLがセットされ何も返されない",
			avatarImageURL: "https://example.com/avatar.png?size=large&format=png",
			wantErr:        nil,
		},
		{
			name:           "異常系: 空のURLの場合、エラーが返される",
			avatarImageURL: "",
			wantErr:        yatter_errors.ErrInternal,
		},
		{
			name:           "異常系: 不正なURL形式の場合、エラーが返される",
			avatarImageURL: "invalid-url",
			wantErr:        yatter_errors.ErrInternal,
		},
		{
			name:           "異常系: スキームがないURLの場合、エラーが返される",
			avatarImageURL: "example.com/avatar.png",
			wantErr:        yatter_errors.ErrInternal,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			p := &profile.Profile{}

			err := p.SetAvatarImageURL(tt.avatarImageURL)

			if !yatter_errors.Is(err, tt.wantErr) {
				t.Errorf("SetAvatarImageURL() error = %v, wantErr %v", err, tt.wantErr)
			}

			// 正常系の場合は、アバター画像URLが正しくセットされているか確認
			if tt.wantErr == nil {
				if p.AvatarImageURL() != tt.avatarImageURL {
					t.Errorf("AvatarImageURL() = %v, want %v", p.AvatarImageURL(), tt.avatarImageURL)
				}
			}
		})
	}
}

func Test_Profile_SetHeaderImageURL(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		headerImageURL string
		wantErr        error
	}{
		{
			name:           "正常系: 正しいURL形式の場合、ヘッダー画像URLがセットされ何も返されない",
			headerImageURL: "https://example.com/header.png",
			wantErr:        nil,
		},
		{
			name:           "正常系: HTTPスキームのURLの場合、ヘッダー画像URLがセットされ何も返されない",
			headerImageURL: "http://example.com/header.jpg",
			wantErr:        nil,
		},
		{
			name:           "正常系: クエリパラメータを含むURLの場合、ヘッダー画像URLがセットされ何も返されない",
			headerImageURL: "https://example.com/header.png?width=1500&height=500",
			wantErr:        nil,
		},
		{
			name:           "異常系: 空のURLの場合、エラーが返される",
			headerImageURL: "",
			wantErr:        yatter_errors.ErrInternal,
		},
		{
			name:           "異常系: 不正なURL形式の場合、エラーが返される",
			headerImageURL: "invalid-url",
			wantErr:        yatter_errors.ErrInternal,
		},
		{
			name:           "異常系: スキームがないURLの場合、エラーが返される",
			headerImageURL: "example.com/header.png",
			wantErr:        yatter_errors.ErrInternal,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			p := &profile.Profile{}

			err := p.SetHeaderImageURL(tt.headerImageURL)

			if !yatter_errors.Is(err, tt.wantErr) {
				t.Errorf("SetHeaderImageURL() error = %v, wantErr %v", err, tt.wantErr)
			}

			// 正常系の場合は、ヘッダー画像URLが正しくセットされているか確認
			if tt.wantErr == nil {
				if p.HeaderImageURL() != tt.headerImageURL {
					t.Errorf("HeaderImageURL() = %v, want %v", p.HeaderImageURL(), tt.headerImageURL)
				}
			}
		})
	}
}

func Test_Profile_SetNote(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		note    string
		wantErr error
	}{
		{
			name:    "正常系: ノートが空文字列の場合、ノートがセットされ何も返されない",
			note:    "",
			wantErr: nil,
		},
		{
			name:    "正常系: ノートが通常の長さの文字列の場合、ノートがセットされ何も返されない",
			note:    "これは私のプロフィールです。",
			wantErr: nil,
		},
		{
			name:    "正常系: ノートに日本語や特殊文字が含まれる場合、ノートがセットされ何も返されない",
			note:    "こんにちは！私は #プログラミング が好きです。\n連絡先: example@example.com",
			wantErr: nil,
		},
		{
			name:    "正常系: ノートが500文字の場合、ノートがセットされ何も返されない",
			note:    strings.Repeat("あ", 500),
			wantErr: nil,
		},
		{
			name:    "異常系: ノートが501文字の場合、エラーが返される",
			note:    strings.Repeat("あ", 501),
			wantErr: yatter_errors.ErrInternal,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			p := &profile.Profile{}

			err := p.SetNote(tt.note)

			if !yatter_errors.Is(err, tt.wantErr) {
				t.Errorf("SetNote() error = %v, wantErr %v", err, tt.wantErr)
			}

			// 正常系の場合は、ノートが正しくセットされているか確認
			if tt.wantErr == nil {
				if p.Note() != tt.note {
					t.Errorf("Note() = %v, want %v", p.Note(), tt.note)
				}
			}
		})
	}
}
