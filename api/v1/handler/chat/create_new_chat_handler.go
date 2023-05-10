package handlerchat

import (
	dtochat "github.com/RyaWcksn/nann-e/dtos/chat"
	customerror "github.com/RyaWcksn/nann-e/pkgs/error"
	"github.com/RyaWcksn/nann-e/pkgs/validator"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (ch *ChatHandler) CreateNewChat(c *fiber.Ctx) error {
	functionName := "RolesHandler.CreateRoles"
	ctx := c.Context()

	payload := &dtochat.CreateNewChatRequest{}
	if err := c.BodyParser(payload); err != nil {
		ch.l.Errorf("[%s - c.BodyParser] : %s", functionName, err)
		return customerror.GetError(customerror.InternalServer, err)
	}

	if err := validator.Validate(payload); err != nil {
		ch.l.Errorf("[%s : validator.Validate] : %s", functionName, err)
		return err
	}

	chatDetail, createChatErr := ch.chatService.CreateNewChat(ctx, payload)
	if createChatErr != nil {
		return createChatErr
	}

	res := struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}{
		Code:    http.StatusCreated,
		Message: http.StatusText(http.StatusCreated),
		Data:    chatDetail,
	}

	return c.Status(fiber.StatusCreated).JSON(res)
}
