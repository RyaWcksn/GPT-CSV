package handlerroles

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (r *RolesHandler) GetOneRoleById(c *fiber.Ctx) error {
	ctx := c.Context()

	roleName := c.Params("roleName")

	roleDetail, getRoleErr := r.rolesService.GetOneRoleById(ctx, roleName)
	if getRoleErr != nil {
		return getRoleErr
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
