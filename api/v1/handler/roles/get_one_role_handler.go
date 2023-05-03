package handlerroles

import (
	dtoroles "github.com/RyaWcksn/nann-e/dtos/roles"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strings"
)

func (r *RolesHandler) GetOneRoleById(c *fiber.Ctx) error {
	ctx := c.Context()

	payload := new(dtoroles.GetOneRoleRequest)
	roleName := c.Params("roleName")
	roleNameClean := strings.ReplaceAll(roleName, "%20", " ")

	payload.RoleName = roleNameClean

	roleDetail, getRoleErr := r.rolesService.GetOneRole(ctx, payload)
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
