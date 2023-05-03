package servicechild

import (
	"context"
	repochild "github.com/RyaWcksn/nann-e/api/v1/repository/user_child"
	dtochild "github.com/RyaWcksn/nann-e/dtos/user_child"
	entitychild "github.com/RyaWcksn/nann-e/entities/user_child"
	"github.com/RyaWcksn/nann-e/pkgs/logger"
)

type ChildService struct {
	childRepo repochild.IRepository
	l         logger.ILogger
}

func NewChildService(childRepo repochild.IRepository, l logger.ILogger) *ChildService {
	return &ChildService{
		childRepo: childRepo,
		l:         l,
	}
}

type IService interface {
	CreateUserChild(ctx context.Context, payload *dtochild.CreateUserChildRequest) (*entitychild.UserChildDetail, error)
	GetOneUserChild(ctx context.Context, payload *dtochild.GetOneUserChildRequest) (*entitychild.UserChildDetail, error)
	GetListUserChild(ctx context.Context, payload *dtochild.GetListUserChildRequest) ([]entitychild.UserChildDetail, error)
	UpdateSingleUserChild(ctx context.Context, payload *dtochild.UpdateSingleUserChildRequest) (*entitychild.UserChildDetail, error)
}

var _ IService = (*ChildService)(nil)
