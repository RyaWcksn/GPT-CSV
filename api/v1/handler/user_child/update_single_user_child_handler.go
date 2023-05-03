package handlerchild

import (
	dtochild "github.com/RyaWcksn/nann-e/dtos/user_child"
	customerror "github.com/RyaWcksn/nann-e/pkgs/error"
	"github.com/RyaWcksn/nann-e/pkgs/validator"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strings"
)

func (ch *ChildHandler) UpdateSingleUserChild(c *fiber.Ctx) error {
	functionName := "ChildHandler.UpdateSingleUserChild"
	ctx := c.Context()

	childName := c.Params("childName")
	childNameClean := strings.ReplaceAll(childName, "%20", " ")

	payload := new(dtochild.UpdateSingleUserChildRequest)
	if err := c.BodyParser(payload); err != nil {
		ch.l.Errorf("[%s - c.BodyParser] : %s", functionName, err)
		return customerror.GetError(customerror.InternalServer, err)
	}

	payload.ChildName = childNameClean

	if err := validator.Validate(payload); err != nil {
		ch.l.Errorf("[%s : validator.Validate] : %s", functionName, err)
		return err
	}

	childDetail, updateChildErr := ch.childService.UpdateSingleUserChild(ctx, payload)
	if updateChildErr != nil {
		return updateChildErr
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
