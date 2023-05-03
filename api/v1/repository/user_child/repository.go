package repochild

import (
	"context"
	dtochild "github.com/RyaWcksn/nann-e/dtos/user_child"
	entitychild "github.com/RyaWcksn/nann-e/entities/user_child"
	storechild "github.com/RyaWcksn/nann-e/store/database/user_child"
)

type IRepository interface {
	CreateUserChild(ctx context.Context, payload *dtochild.CreateUserChildRequest) error
	GetOneUserChild(ctx context.Context, payload *dtochild.GetOneUserChildRequest) (*entitychild.UserChildDetail, error)
	GetListUserChild(ctx context.Context, offset, limit int) ([]entitychild.UserChildDetail, error)
	UpdateSingleUserChild(ctx context.Context, payload *dtochild.UpdateSingleUserChildRequest) error
}

var _ IRepository = (*storechild.ChildImpl)(nil)
