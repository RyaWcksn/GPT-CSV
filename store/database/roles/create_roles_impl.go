package storeroles

import (
	"context"
	dtoroles "github.com/RyaWcksn/nann-e/dtos/roles"
	customerror "github.com/RyaWcksn/nann-e/pkgs/error"
)

func (r *RolesImpl) CreateRoles(ctx context.Context, payload *dtoroles.CreateRoleRequest) error {
	functionName := "RolesImpl.CreateRoles"

	tx, err := r.DB.Begin()
	if err != nil {
		r.l.Errorf("[%s : o.DB.Begin - error begin transaction] : %v", functionName, err)
		return customerror.GetError(customerror.InternalServer, err)
	}

	stmt, err := tx.PrepareContext(ctx, QueryCreateRoles)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			r.l.Errorf("update failed: %v, unable to back: %v", err, rollbackErr)
			return customerror.GetError(customerror.InternalServer, err)
		}
		r.l.Errorf("[%s - tx.PrepareContext()] : %s", functionName, err)
		return customerror.GetError(customerror.InternalServer, err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		payload.RoleId,
		payload.Topic,
		payload.ChildDescription,
		payload.RoleName,
		payload.RoleDescription,
	)

	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			r.l.Errorf("update failed: %v, unable to back: %v", err, rollbackErr)
			return customerror.GetError(customerror.InternalServer, err)
		}

		r.l.Errorf("[%s - stmt.ExecContext]: %s", functionName, err)
		return customerror.GetError(customerror.InternalServer, err)
	}

	if commitErr := tx.Commit(); commitErr != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			r.l.Errorf("update failed: %v, unable to back: %v", err, rollbackErr)
			return customerror.GetError(customerror.InternalServer, err)
		}
		r.l.Errorf("update failed: %v, unable to commit: %v", err, commitErr)
		return customerror.GetError(customerror.InternalServer, err)
	}

	return nil
}

