package repousersparent

import (
	"context"
	dtoauthentication "github.com/RyaWcksn/nann-e/dtos/authentication"
	entityauthentication "github.com/RyaWcksn/nann-e/entities/authentication"

	storeusersparent "github.com/RyaWcksn/nann-e/store/database/user"
)

type IRepository interface {
	CreateUsersParent(ctx context.Context, payload *dtoauthentication.RegisterRequest) error
	GetOneUsersParentById(ctx context.Context, id string) (*entityauthentication.RegisterDetails, error)
}

var _ IRepository = (*storeusersparent.UserParentImpl)(nil)
