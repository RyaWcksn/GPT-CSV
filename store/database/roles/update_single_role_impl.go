package storeroles

import (
	"context"
	dtoroles "github.com/RyaWcksn/nann-e/dtos/roles"
	customerror "github.com/RyaWcksn/nann-e/pkgs/error"
)

func (r *RolesImpl) UpdateSingleRoleById(ctx context.Context, payload *dtoroles.UpdateSingleRoleRequest) error {
	functionName := "RolesImpl.UpdateSingleRoleById"

	tx, err := r.DB.Begin()
	if err != nil {
		r.l.Errorf("[%s : o.DB.Begin - error begin transaction] : %v", functionName, err)
		return customerror.GetError(customerror.InternalServer, err)
	}

	_, err = tx.ExecContext(ctx, QueryUpdateSingleRoleById,
		payload.Topic,
		payload.Rules,
		payload.Goals,
		payload.ChildDescription,
		payload.RoleDescription,
		payload.ParentId,
		payload.RoleName,
	)
	if err != nil {
		r.l.Debugf("[%s : tx.ExecContext] : %s", functionName, err)
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			r.l.Errorf("[%s : tx.Rollback] : %s", functionName, rollbackErr)
			return customerror.GetError(customerror.InternalServer, rollbackErr)
		}
		return customerror.GetError(customerror.InternalServer, err)
	}

	if commitErr := tx.Commit(); commitErr != nil {
		r.l.Errorf("[%s : tx.Commit] : %s", functionName, commitErr)
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			r.l.Errorf("[%s : tx.Rollback] : %s", functionName, rollbackErr)
			return customerror.GetError(customerror.InternalServer, rollbackErr)
		}
		return customerror.GetError(customerror.InternalServer, commitErr)
	}

	return nil
}
