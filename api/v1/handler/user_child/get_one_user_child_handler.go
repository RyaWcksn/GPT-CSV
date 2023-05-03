package handlerchild

import (
	dtochild "github.com/RyaWcksn/nann-e/dtos/user_child"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strings"
)

func (ch *ChildHandler) GetOneUserChild(c *fiber.Ctx) error {
	ctx := c.Context()

	payload := new(dtochild.GetOneUserChildRequest)
	childName := c.Params("childName")
	childNameClean := strings.ReplaceAll(childName, "%20", " ")

	payload.ChildName = childNameClean

	childDetail, getChildErr := ch.childService.GetOneUserChild(ctx, payload)
	if getChildErr != nil {
		return getChildErr
	}

	res := struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    childDetail,
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
