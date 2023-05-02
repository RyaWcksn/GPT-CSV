package middleware

import (
	"errors"
	customerror "github.com/RyaWcksn/nann-e/pkgs/error"
	"github.com/gofiber/fiber/v2"
	"strings"
)

// ErrorHandler is a middleware that write any error that occurred to the server
func ErrorHandler(c *fiber.Ctx) error {
	err := c.Next()
	if err != nil {
		if strings.HasPrefix(err.Error(), "Method Not Allowed") || strings.HasPrefix(err.Error(), "Cannot") {
			xerr := customerror.GetError(customerror.RequestNotFound, errors.New("invalid endpoint or using wrong method"))
			return c.Status(xerr.Code).JSON(xerr)
		}
		xerr, ok := err.(*customerror.ErrorForm)
		if !ok {
			xerr = customerror.GetError(customerror.InternalServer, errors.New(customerror.InternalServer))
		}
		return c.Status(xerr.Code).JSON(xerr)
	}

	return nil
}
