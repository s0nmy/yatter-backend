package query

import (
	"context"
	"database/sql"
	stderr "errors"
	"yatter-backend-go/app/domain/object/auth"
	"yatter-backend-go/app/usecase/query"
	"yatter-backend-go/pkg/errors"

	"github.com/jmoiron/sqlx"
)

var _ query.AuthQueryService = (*AuthQueryServiceImpl)(nil)

type AuthQueryServiceImpl struct {
	db *sqlx.DB
}

func NewAuthQueryService(db *sqlx.DB) *AuthQueryServiceImpl {
	return &AuthQueryServiceImpl{db: db}
}

type FindCredentialByUsernameDTO struct {
	Username     string `db:"username"`
	PasswordHash string `db:"password_hash"`
}

func (s *AuthQueryServiceImpl) FindCredentialByUsername(
	ctx context.Context,
	username string,
) (*auth.Credential, error) {
	var dbCredential FindCredentialByUsernameDTO
	err := s.db.GetContext(ctx, &dbCredential,
		`SELECT username, password_hash FROM user WHERE username = ?`,
		username,
	)

	if err != nil {
		if stderr.Is(err, sql.ErrNoRows) {
			return nil, errors.ErrBadRequest.WithDevMessage("user not found")
		}
		return nil, err
	}

	credential, err := auth.ReconstructCredential(dbCredential.Username, dbCredential.PasswordHash)
	if err != nil {
		return nil, err
	}

	return credential, nil
}
