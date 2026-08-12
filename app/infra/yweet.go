package infra

import (
	"context"
	"time"
	"yatter-backend-go/app/domain/object/yweet"
	"yatter-backend-go/app/domain/repository"
	"yatter-backend-go/app/infra/transaction"
)

// repoで書いた関数があることの確認するための定義
var _ repository.Yweet = (*YweetRepoImpl)(nil)

// app/domain/repository/yweet.go にてインターフェースは定義しているため、ここでは行わない
type YweetRepoImpl struct{}

// server.go で認識できるようにするための
func NewYweetRepository() *YweetRepoImpl {
	return &YweetRepoImpl{}
}

// yweetDTO: ユーザー用のデータ詰め替え構造体
// DBからデータ取得する際の、テーブル定義 <-> ドメインモデルの変換を行う
// TODO: 本当は別ファイルに定義した方がわかりやすそう
type insertedYweetDTO struct {
	ID        uint64    `db:"id"`
	UserID    uint64    `db:"user_id"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
}

func (y *YweetRepoImpl) Insert(ctx context.Context, Yweet *yweet.Yweet) (*yweet.Yweet, error) {
	// トランザクションを張る
	tx, err := transaction.FetchTransaction(ctx)
	if err != nil {
		return nil, err
	}

	insertResult, err := tx.ExecContext(
		ctx,
		`INSERT INTO yweet (user_id, content) VALUES (?, ?)`,
		Yweet.UserID(),
		Yweet.Content(),
	)
	if err != nil {
		return nil, err
	}

	yweetID, err := insertResult.LastInsertId()
	if err != nil {
		return nil, err
	}

	var insertedYweetDTO insertedYweetDTO
	err = tx.GetContext(ctx, &insertedYweetDTO, `SELECT id, user_id, content, created_at FROM yweet WHERE id = ?`, yweetID)
	if err != nil {
		// NOTE: インサート済みであるはずなので、 sql.NoRows の場合でもエラーとして返す
		return nil, err
	}

	Yweet.SetID(insertedYweetDTO.ID)
	Yweet.SetUserID(insertedYweetDTO.UserID)
	Yweet.SetContent(insertedYweetDTO.Content)
	Yweet.SetCreatedAt(insertedYweetDTO.CreatedAt)

	return Yweet, nil
}
