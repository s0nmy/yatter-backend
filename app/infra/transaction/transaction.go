package transaction

import (
	"context"
	"log/slog"
	"yatter-backend-go/app/usecase/transactor"
	"yatter-backend-go/pkg/errors"

	"github.com/jmoiron/sqlx"
)

var _ transactor.Transactor = (*TransactorImpl)(nil)

type TransactorImpl struct {
	db *sqlx.DB
}

func NewTransactor(db *sqlx.DB) *TransactorImpl {
	return &TransactorImpl{
		db: db,
	}
}

type transactionContextKey struct{}

var transactionKey = transactionContextKey{}

func (t *TransactorImpl) Transaction(ctx context.Context, txFunc func(ctx context.Context) error) error {
	tx, err := t.db.Beginx()
	if err != nil {
		return err
	}

	ctxWithTx := context.WithValue(ctx, transactionKey, tx)

	err = txFunc(ctxWithTx)
	if err != nil {
		return tx.Rollback()
	}

	return tx.Commit()
}

func (t *TransactorImpl) TransactionWithValue(
	ctx context.Context,
	txFunc func(ctx context.Context) (any, error),
) (any, error) {
	tx, err := t.db.Beginx()
	if err != nil {
		return nil, err
	}

	ctxWithTx := context.WithValue(ctx, transactionKey, tx)

	result, err := txFunc(ctxWithTx)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			// ロールバックに失敗した場合はエラーログを出力するのみで、エラーは返さない
			// ロールバックの失敗を返してもアプリケーション側で取れる手段がないため
			slog.ErrorContext(ctx, "failed to rollback", "err", rollbackErr)
		}
		return nil, err
	}

	return result, tx.Commit()
}

func FetchTransaction(ctx context.Context) (*sqlx.Tx, error) {
	tx, ok := ctx.Value(transactionKey).(*sqlx.Tx)
	if !ok {
		return nil, errors.ErrInternal.WithDevMessage("transaction not found in context")
	}

	return tx, nil
}
