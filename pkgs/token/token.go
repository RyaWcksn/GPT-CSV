package tokens

import (
	dtoauthentication "github.com/RyaWcksn/nann-e/dtos/authentication"
	"github.com/golang-jwt/jwt"
	"time"
)

func GenerateJWT(payload *dtoauthentication.TokenRequest) (string, error) {
	// create a new token
	token := jwt.New(jwt.SigningMethodHS256)

	// set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = payload.Id
	claims["exp"] = time.Now().Add(30 * time.Minute).Unix()

	// generate encoded token and return
	return token.SignedString([]byte(payload.SecretKey))
}
