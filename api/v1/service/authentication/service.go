package serviceauthentication

import (
	"context"
	repousersparent "github.com/RyaWcksn/nann-e/api/v1/repository/users_parent"
	"github.com/RyaWcksn/nann-e/config"
	dtoauthentication "github.com/RyaWcksn/nann-e/dtos/authentication"
	entityauthentication "github.com/RyaWcksn/nann-e/entities/authentication"
	"github.com/RyaWcksn/nann-e/pkgs/logger"
)

type AuthenticationService struct {
	usersParentRepo repousersparent.IRepository
	cfg             *config.Config
	l               logger.ILogger
}

func NewServiceImpl(usersParentRepo repousersparent.IRepository, cfg *config.Config, log logger.ILogger) *AuthenticationService {
	return &AuthenticationService{
		usersParentRepo: usersParentRepo,
		cfg:             cfg,
		l:               log,
	}
}

type IService interface {
	RegisterParent(ctx context.Context, payload *dtoauthentication.RegisterRequest) (*entityauthentication.RegisterDetails, error)
	LoginParent(ctx context.Context, payload *dtoauthentication.LoginRequest) (*entityauthentication.LoginDetails, error)
}

var _ IService = (*AuthenticationService)(nil)
