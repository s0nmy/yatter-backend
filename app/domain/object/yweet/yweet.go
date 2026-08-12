package yweet

import (
	"time"
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

type Yweet struct {
	id        uint64
	userID    uint64
	content   string
	createdAt time.Time
}

// サーバー側で決めない値を返却
func NewYweet(id uint64, userID uint64, content string, createdAt time.Time) (*Yweet, error) {
	// 空のインスタンスを作成
	yweet := &Yweet{}

	yweet.SetID(id)
	yweet.SetUserID(userID)
	yweet.SetContent(content)
	yweet.SetCreatedAt(createdAt)

	return yweet, nil

}

func (y *Yweet) SetID(id uint64) {
	y.id = id
}

func (y *Yweet) ID() uint64 {
	return y.id
}

func (y *Yweet) SetUserID(userID uint64) {
	y.userID = userID
}

func (y *Yweet) UserID() uint64 {
	return y.userID
}

func (y *Yweet) SetContent(content string) {
	y.content = content
}

func (y *Yweet) Content() string {
	return y.content
}

func (y *Yweet) SetCreatedAt(createdAt time.Time) {
	y.createdAt = createdAt
}

func (y *Yweet) CreatedAt() time.Time {
	return y.createdAt
}
