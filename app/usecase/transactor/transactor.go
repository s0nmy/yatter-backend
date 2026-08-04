package transactor

import "context"

type Transactor interface {
	Transaction(ctx context.Context, txFunc func(context.Context) error) error

	// TransactionWithValue はトランザクション内で値を返す場合に利用する
	TransactionWithValue(ctx context.Context, txFunc func(context.Context) (any, error)) (any, error)
}
