package handlerroles

import (
	dtoroles "github.com/RyaWcksn/nann-e/dtos/roles"
	"github.com/RyaWcksn/nann-e/pkgs/validator"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

func (r *RolesHandler) GetListRole(c *fiber.Ctx) error {
	functionName := "ROlesHandler.GetListRole"
	ctx := c.Context()

	pageNumber := c.Query("page")
	limit := c.Query("limit")

	payload := new(dtoroles.GetListRoleRequest)
	payload.PageNumber, _ = strconv.Atoi(pageNumber)
	payload.Limit, _ = strconv.Atoi(limit)

	if err := validator.Validate(payload); err != nil {
		r.l.Errorf("[%s : validator.Validate] : %s", functionName, err)
		return err
	}

	roleDetail, getListErr := r.rolesService.GetListRole(ctx, payload)
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
		Data:    roleDetail,
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

