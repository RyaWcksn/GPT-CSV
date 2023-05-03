package handlerchild

import (
	dtochild "github.com/RyaWcksn/nann-e/dtos/user_child"
	customerror "github.com/RyaWcksn/nann-e/pkgs/error"
	"github.com/RyaWcksn/nann-e/pkgs/validator"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (ch *ChildHandler) CreateUserChild(c *fiber.Ctx) error {
	functionName := "RolesHandler.CreateRoles"
	ctx := c.Context()

	payload := &dtochild.CreateUserChildRequest{}
	if err := c.BodyParser(payload); err != nil {
		ch.l.Errorf("[%s - c.BodyParser] : %s", functionName, err)
		return customerror.GetError(customerror.InternalServer, err)
	}

	if err := validator.Validate(payload); err != nil {
		ch.l.Errorf("[%s : validator.Validate] : %s", functionName, err)
		return err
	}

	childDetail, createUserChildErr := ch.childService.CreateUserChild(ctx, payload)
	if createUserChildErr != nil {
		return createUserChildErr
	}

	res := struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}{
		Code:    http.StatusCreated,
		Message: http.StatusText(http.StatusCreated),
		Data:    childDetail,
	}

	return c.Status(fiber.StatusCreated).JSON(res)
}
