package handlerusersparent

import (
	serviceauthentication "github.com/RyaWcksn/nann-e/api/v1/service/authentication"
	"github.com/RyaWcksn/nann-e/pkgs/logger"
	"github.com/gofiber/fiber/v2"
)

type AuthenticationHandler struct {
	authService serviceauthentication.IService
	l           logger.ILogger
}

func NewUsersParentHandler(usersParentService serviceauthentication.IService, l logger.ILogger) *AuthenticationHandler {
	return &AuthenticationHandler{
		authService: usersParentService,
		l:           l,
	}
}

type IHandler interface {
	RegisterParent(c *fiber.Ctx) error
	LoginParent(c *fiber.Ctx) error
}

var _ IHandler = (*AuthenticationHandler)(nil)
