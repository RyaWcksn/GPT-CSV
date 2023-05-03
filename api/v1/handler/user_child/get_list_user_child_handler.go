package handlerchild

import (
	dtochild "github.com/RyaWcksn/nann-e/dtos/user_child"
	"github.com/RyaWcksn/nann-e/pkgs/validator"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

func (ch *ChildHandler) GetListUserChild(c *fiber.Ctx) error {
	functionName := "ChildHandler.GetListUserChild"
	ctx := c.Context()

	pageNumber := c.Query("page")
	limit := c.Query("limit")

	payload := new(dtochild.GetListUserChildRequest)
	payload.PageNumber, _ = strconv.Atoi(pageNumber)
	payload.Limit, _ = strconv.Atoi(limit)

	if err := validator.Validate(payload); err != nil {
		ch.l.Errorf("[%s : validator.Validate] : %s", functionName, err)
		return err
	}

	childDetail, getListErr := ch.childService.GetListUserChild(ctx, payload)
	if getListErr != nil {
		return getListErr
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

