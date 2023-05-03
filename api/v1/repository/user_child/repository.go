package repochild

import (
	"context"
	dtochild "github.com/RyaWcksn/nann-e/dtos/user_child"
	storechild "github.com/RyaWcksn/nann-e/store/database/user_child"
)

type IRepository interface {
	CreateUserChild(ctx context.Context, payload *dtochild.CreateUserChildRequest) error
}

var _ IRepository = (*storechild.ChildImpl)(nil)
