package serviceauthentication

import (
	"context"
	dtoauthentication "github.com/RyaWcksn/nann-e/dtos/authentication"
	entityauthentication "github.com/RyaWcksn/nann-e/entities/authentication"
	"github.com/RyaWcksn/nann-e/pkgs/utils"

	customerror "github.com/RyaWcksn/nann-e/pkgs/error"
	"golang.org/x/crypto/bcrypt"
)

func (u *AuthenticationService) RegisterParent(ctx context.Context, payload *dtoauthentication.RegisterRequest) (*entityauthentication.RegisterDetails, error) {
	functionName := "UsersParentService.RegisterParent"

	// Generate users parent id
	parentId, generateErr := utils.GenerateUUIDFromEmailAndPhoneNumber(payload.Email, payload.PhoneNumber)
	if generateErr != nil {
		u.l.Errorf("[%s : utils.GenerateUUIDFromEmailAndPhoneNumber] : %s", functionName, generateErr)
		return nil, customerror.GetError(customerror.BadRequest, generateErr)
	}

	payload.UsersParentId = parentId

	// hash password
	hashPassword, errHashPass := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if errHashPass != nil {
		u.l.Errorf("[%s : bcrypt.GenerateFromPassword] : %s", functionName, errHashPass)
		return nil, errHashPass
	}
	payload.Password = string(hashPassword)

	payload.Status = 1

	// call repo func
	err := u.usersParentRepo.CreateUsersParent(ctx, payload)
	if err != nil {
		u.l.Errorf("[%s : u.usersParentRepo.CreateUsersParent] : %s", functionName, err)
		return nil, err
	}

	res := entityauthentication.RegisterDetails{
		UsersParentId: payload.UsersParentId,
		Name:          payload.Name,
		Email:         payload.Email,
		PhoneNumber:   payload.PhoneNumber,
		Status:        payload.Status,
	}

	return &res, nil
}
