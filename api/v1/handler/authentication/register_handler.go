package handlerusersparent

import (
	dtoauthentication "github.com/RyaWcksn/nann-e/dtos/authentication"
	customerror "github.com/RyaWcksn/nann-e/pkgs/error"
	"github.com/RyaWcksn/nann-e/pkgs/validator"
	"github.com/gofiber/fiber/v2"
)

func (u *AuthenticationHandler) RegisterParent(c *fiber.Ctx) error {
	functionName := "UsersParentHandler.RegisterParent"
	ctx := c.Context()

	payload := &dtoauthentication.RegisterRequest{}
	if err := c.BodyParser(payload); err != nil {
		u.l.Errorf("[%s - c.BodyParser] : %s", functionName, err)
		return customerror.GetError(customerror.InternalServer, err)
	}

	if err := validator.Validate(payload); err != nil {
		u.l.Errorf("[%s : validator.Validate] : %s", functionName, err)
		return err
	}

	registerDetails, registerErr := u.authService.RegisterParent(ctx, payload)
	if registerErr != nil {
		u.l.Errorf("[%s : u.usersParentService.CreateUserParent] : %s", functionName, registerErr)
		return registerErr
	}

	return c.Status(fiber.StatusCreated).JSON(registerDetails)
}
