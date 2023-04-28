package handlerusersparent

import (
	dtoauthentication "github.com/RyaWcksn/nann-e/dtos/authentication"
	customerror "github.com/RyaWcksn/nann-e/pkgs/error"
	"github.com/RyaWcksn/nann-e/pkgs/validator"
	"github.com/gofiber/fiber/v2"
)

func (u *AuthenticationHandler) LoginParent(c *fiber.Ctx) error {
	functionName := "UsersParentHandler.LoginParent"
	ctx := c.Context()

	payload := &dtoauthentication.LoginRequest{}
	if err := c.BodyParser(payload); err != nil {
		u.l.Errorf("[%s - c.BodyParser] : %s", functionName, err)
		return customerror.GetError(customerror.InternalServer, err)
	}

	if err := validator.Validate(payload); err != nil {
		u.l.Errorf("[%s : validator.Validate] : %s", functionName, err)
		return err
	}

	loginDetails, loginErr := u.authService.LoginParent(ctx, payload)
	if loginErr != nil {
		u.l.Errorf("[%s : u.authService.LoginParent] : %s", functionName, loginErr)
		return loginErr
	}

	return c.Status(fiber.StatusOK).JSON(loginDetails)
}
