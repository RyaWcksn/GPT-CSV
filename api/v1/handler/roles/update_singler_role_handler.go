package handlerroles

import (
	dtoroles "github.com/RyaWcksn/nann-e/dtos/roles"
	customerror "github.com/RyaWcksn/nann-e/pkgs/error"
	"github.com/RyaWcksn/nann-e/pkgs/validator"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strings"
)

func (r *RolesHandler) UpdateSingleRole(c *fiber.Ctx) error {
	functionName := "RolesHandler.UpdateSingleRole"
	ctx := c.Context()

	roleName := c.Params("roleName")
	roleNameClean := strings.ReplaceAll(roleName, "%20", " ")

	payload := &dtoroles.UpdateSingleRoleRequest{}
	if err := c.BodyParser(payload); err != nil {
		r.l.Errorf("[%s - c.BodyParser] : %s", functionName, err)
		return customerror.GetError(customerror.InternalServer, err)
	}
	payload.RoleName = roleNameClean

	if err := validator.Validate(payload); err != nil {
		r.l.Errorf("[%s : validator.Validate] : %s", functionName, err)
		return err
	}

	roleDetail, updateRoleErr := r.rolesService.UpdateSingleRoleById(ctx, payload)
	if updateRoleErr != nil {
		return updateRoleErr
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
