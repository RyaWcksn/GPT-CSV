package middleware

import (
	"errors"
	"github.com/RyaWcksn/nann-e/config"
	customerror "github.com/RyaWcksn/nann-e/pkgs/error"
	"github.com/RyaWcksn/nann-e/pkgs/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"strings"
)

func Authenticate(cfg *config.Config, l logger.ILogger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := c.Get("Authorization")

		if tokenString == "" {
			xerr := customerror.GetError(customerror.Unauthorized, errors.New("token is invalid or expired"))
			return c.Status(xerr.Code).JSON(xerr)
		}

		// remove "Bearer " prefix from token
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Verify the token signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrInvalidKey
			}
			// Return the secret key used to sign the token
			return []byte(cfg.App.SECRET), nil
		})

		if err != nil {
			l.Errorf("[MIDDLEWARE AUTHENTICATION - error retrieving token]")
			xerr := customerror.GetError(customerror.Unauthorized, errors.New("token is invalid or expired"))
			return c.Status(xerr.Code).JSON(xerr)
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			l.Errorf("[MIDDLEWARE AUTHENTICATION - error claims token")
			xerr := customerror.GetError(customerror.Unauthorized, errors.New("token is invalid or expired"))
			return c.Status(xerr.Code).JSON(xerr)
		}

		id := claims["id"].(string)
		exp := claims["exp"].(float64)
		c.Locals("ctxParentId", id)
		c.Locals("ctxExpireTime", exp)

		return c.Next()
	}
}
