package serviceauthentication

import (
	"context"
	dtoauthentication "github.com/RyaWcksn/nann-e/dtos/authentication"
	entityauthentication "github.com/RyaWcksn/nann-e/entities/authentication"
	customerror "github.com/RyaWcksn/nann-e/pkgs/error"
	tokens "github.com/RyaWcksn/nann-e/pkgs/token"
	"github.com/RyaWcksn/nann-e/pkgs/utils"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func (u *AuthenticationService) LoginParent(ctx context.Context, payload *dtoauthentication.LoginRequest) (*entityauthentication.LoginDetails, error) {
	functionName := "UsersParentService.LoginParent"

	parentId, generateErr := utils.GenerateUUIDFromEmailAndPhoneNumber(payload.Email, payload.PhoneNumber)
	if generateErr != nil {
		u.l.Errorf("[%s : utils.GenerateUUIDFromEmailAndPhoneNumber] : %s", functionName, generateErr)
		return nil, customerror.GetError(customerror.BadRequest, generateErr)
	}

	parentDetail, getOneUsersParentErr := u.usersParentRepo.GetOneUsersParentById(ctx, parentId)
	if getOneUsersParentErr != nil {
		u.l.Errorf("%s : u.usersParentRepo.GetOneUsersParentById", functionName, getOneUsersParentErr)
		return nil, getOneUsersParentErr
	}

	// check password
	if err := bcrypt.CompareHashAndPassword([]byte(parentDetail.Password), []byte(payload.Password)); err != nil {
		u.l.Errorf("[%s : bcrypt.CompareHashAndPassword - password missmatch] : %s", functionName, err)
		return nil, customerror.GetError(customerror.BadRequest, err)
	}

	tokenRequest := &dtoauthentication.TokenRequest{
		Id:        parentId,
		SecretKey: u.cfg.App.SECRET,
	}

	token, err := tokens.GenerateJWT(tokenRequest)
	if err != nil {
		u.l.Errorf("[%s : tokens.GenerateJWT - error generating JWT token] : %s", functionName, err)
		return nil, customerror.GetError(customerror.InternalServer, err)
	}

	res := entityauthentication.LoginDetails{
		Token:      token,
		ParentId:   parentId,
		ExpiryDate: time.Now().Add(30 * time.Minute).Unix(),
	}

	return &res, nil
}
