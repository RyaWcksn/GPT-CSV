package handlerchild

import (
	servicechild "github.com/RyaWcksn/nann-e/api/v1/service/user_child"
	"github.com/RyaWcksn/nann-e/pkgs/logger"
	"github.com/gofiber/fiber/v2"
)

type ChildHandler struct {
	childService servicechild.IService
	l            logger.ILogger
}

func NewChildHandler(childService servicechild.IService, l logger.ILogger) *ChildHandler {
	return &ChildHandler{
		childService: childService,
		l:            l,
	}
}

type IHandler interface {
	CreateUserChild(c *fiber.Ctx) error
}
