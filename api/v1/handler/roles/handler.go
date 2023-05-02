package handlerroles

import (
	serviceroles "github.com/RyaWcksn/nann-e/api/v1/service/roles"
	"github.com/RyaWcksn/nann-e/pkgs/logger"
	"github.com/gofiber/fiber/v2"
)

type RolesHandler struct {
	rolesService serviceroles.IService
	l            logger.ILogger
}

func NewRoles(rolesService serviceroles.IService, l logger.ILogger) *RolesHandler {
	return &RolesHandler{
		rolesService: rolesService,
		l:            l,
	}
}

type IHandler interface {
	CreateRoles(c *fiber.Ctx) error
	GetOneRoleById(c *fiber.Ctx) error
	GetListRole(c *fiber.Ctx) error
}

var _ IHandler = (*RolesHandler)(nil)
