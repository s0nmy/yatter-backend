package query

import (
	"context"
	"yatter-backend-go/app/domain/object/auth"
)

type AuthQueryService interface {
	FindCredentialByUsername(ctx context.Context, username string) (*auth.Credential, error)
}
