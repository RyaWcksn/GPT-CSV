package handlerroles

import (
	dtoroles "github.com/RyaWcksn/nann-e/dtos/roles"
	customerror "github.com/RyaWcksn/nann-e/pkgs/error"
	"github.com/RyaWcksn/nann-e/pkgs/validator"
	"github.com/gofiber/fiber/v2"
)

func (r *RolesHandler) CreateRoles(c *fiber.Ctx) error {
	functionName := "RolesHandler.CreateRoles"
	ctx := c.Context()

	payload := &dtoroles.CreateRoleRequest{}
	if err := c.BodyParser(payload); err != nil {
		r.l.Errorf("[%s - c.BodyParser] : %s", functionName, err)
		return customerror.GetError(customerror.InternalServer, err)
	}

	if err := validator.Validate(payload); err != nil {
		r.l.Errorf("[%s : validator.Validate] : %s", functionName, err)
		return err
	}

	rolesDetail, createRolesErr := r.rolesService.CreateRoles(ctx, payload)
	if createRolesErr != nil {
		return createRolesErr
	}

	return c.Status(fiber.StatusOK).JSON(rolesDetail)

}

