package storeusersparent

import (
	"context"
	dtoauthentication "github.com/RyaWcksn/nann-e/dtos/authentication"
	customerror "github.com/RyaWcksn/nann-e/pkgs/error"
)

func (u *UserParentImpl) CreateUsersParent(ctx context.Context, payload *dtoauthentication.RegisterRequest) error {
	functionName := "UserParentImpl.CreateUsersParent"
	tx, err := u.DB.Begin()
	if err != nil {
		u.l.Errorf("[%s : o.DB.Begin - error begin transaction] : %v", functionName, err)
		return customerror.GetError(customerror.InternalServer, err)
	}

	stmt, err := tx.PrepareContext(ctx, QuerySaveUsersParent)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			u.l.Errorf("update failed: %v, unable to back: %v", err, rollbackErr)
			return customerror.GetError(customerror.InternalServer, err)
		}
		u.l.Errorf("[%s - tx.PrepareContext()] : %s", functionName, err)
		return customerror.GetError(customerror.InternalServer, err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		payload.UsersParentId,
		payload.Password,
		payload.Name,
		payload.Email,
		payload.PhoneNumber,
		payload.Status,
	)

	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			u.l.Errorf("update failed: %v, unable to back: %v", err, rollbackErr)
			return customerror.GetError(customerror.InternalServer, err)
		}

		u.l.Errorf("[%s - stmt.ExecContext]: %s", functionName, err)
		return customerror.GetError(customerror.InternalServer, err)
	}

	if commitErr := tx.Commit(); commitErr != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			u.l.Errorf("update failed: %v, unable to back: %v", err, rollbackErr)
			return customerror.GetError(customerror.InternalServer, err)
		}
		u.l.Errorf("update failed: %v, unable to commit: %v", err, commitErr)
		return customerror.GetError(customerror.InternalServer, err)
	}

	return nil
}
