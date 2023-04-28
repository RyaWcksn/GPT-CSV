package storeusersparent

import (
	"context"
	"encoding/json"
	entityauthentication "github.com/RyaWcksn/nann-e/entities/authentication"
	customerror "github.com/RyaWcksn/nann-e/pkgs/error"
)

func (u *UserParentImpl) GetOneUsersParentById(ctx context.Context, id string) (*entityauthentication.RegisterDetails, error) {
	functionName := "UserParentImpl.GetOneUsersParentById"
	res := &entityauthentication.RegisterDetails{}

	err := u.DB.QueryRowContext(ctx, QuerySelectOneUsersParentById, id).Scan(
		&res.Password,
		&res.Name,
		&res.Email,
		&res.PhoneNumber,
		&res.Status,
	)
	if err != nil {
		u.l.Errorf("[%s : u.DB.QueryRowContext]", functionName, err)
		return nil, customerror.GetError(customerror.InternalServer, err)
	}

	resByte, _ := json.Marshal(res)
	u.l.Debugf("Query Result : %s", resByte)

	return res, nil
}
