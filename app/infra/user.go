package infra

import (
	"context"
	"database/sql"
	"errors"
	"time"
	"yatter-backend-go/app/domain/object/user"
	"yatter-backend-go/app/domain/repository"
	"yatter-backend-go/app/infra/transaction"
)

var _ repository.User = (*UserRepoImpl)(nil)

type UserRepoImpl struct{}

func NewUserRepository() *UserRepoImpl {
	return &UserRepoImpl{}
}

// userDTO: ユーザー用のデータ詰め替え構造体
// DBからデータ取得する際の、テーブル定義 <-> ドメインモデルの変換を行う
// TODO: 本当は別ファイルに定義した方がわかりやすそう
type userDTO struct {
	ID           uint64    `db:"id"`
	Username     string    `db:"username"`
	PasswordHash string    `db:"password_hash"`
	CreatedAt    time.Time `db:"created_at"`
}

func (ur *UserRepoImpl) FindByUsername(ctx context.Context, username string) (*user.User, error) {
	tx, err := transaction.FetchTransaction(ctx)
	if err != nil {
		return nil, err
	}

	var dbUser userDTO
	err = tx.GetContext(ctx, &dbUser, `SELECT id, username, password_hash, created_at FROM user WHERE username = ?`, username)
	if err != nil {
		//nolint: nilnil // ユーザーが存在しない場合は nil, nil を返す
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	usr, err := user.ReconstructUser(dbUser.ID, dbUser.Username, dbUser.PasswordHash, dbUser.CreatedAt)
	if err != nil {
		return nil, err
	}

	return usr, nil
}

type insertedUserDTO struct {
	ID           uint64    `db:"id"`
	Username     string    `db:"username"`
	PasswordHash string    `db:"password_hash"`
	CreatedAt    time.Time `db:"created_at"`
}

func (ur *UserRepoImpl) Insert(ctx context.Context, pendingUser *user.PendingUser) (*user.User, error) {
	tx, err := transaction.FetchTransaction(ctx)
	if err != nil {
		return nil, err
	}

	insertResult, err := tx.ExecContext(
		ctx,
		`INSERT INTO user (username, password_hash, display_name, avatar, header, note) VALUES (?, ?, ?, ?, ?, ?)`,
		pendingUser.Username(),
		pendingUser.PasswordHash(),
		"",
		"",
		"",
		"",
	)
	if err != nil {
		return nil, err
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		return nil, err
	}

	var insertedUserDTO insertedUserDTO
	err = tx.GetContext(ctx, &insertedUserDTO, `SELECT id, username, password_hash, created_at FROM user WHERE id = ?`, userID)
	if err != nil {
		// NOTE: インサート済みであるはずなので、 sql.NoRows の場合でもエラーとして返す
		return nil, err
	}

	usr, err := user.ReconstructUser(uint64(userID), insertedUserDTO.Username, insertedUserDTO.PasswordHash, insertedUserDTO.CreatedAt)
	if err != nil {
		return nil, err
	}

	return usr, nil
}
