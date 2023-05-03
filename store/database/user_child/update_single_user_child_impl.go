package storechild

import (
	"context"
	dtochild "github.com/RyaWcksn/nann-e/dtos/user_child"
	customerror "github.com/RyaWcksn/nann-e/pkgs/error"
)

func (c *ChildImpl) UpdateSingleUserChild(ctx context.Context, payload *dtochild.UpdateSingleUserChildRequest) error {
	functionName := "ChildImpl.UpdateSingleUserChild"

	tx, err := c.DB.Begin()
	if err != nil {
		c.l.Errorf("[%s : o.DB.Begin - error begin transaction] : %v", functionName, err)
		return customerror.GetError(customerror.InternalServer, err)
	}

	_, err = tx.ExecContext(ctx, QueryUpdateSingleChild,
		payload.RoleName,
		payload.Age,
		payload.ParentId,
		payload.ChildName,
	)
	if err != nil {
		c.l.Debugf("[%s : tx.ExecContext] : %s", functionName, err)
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			c.l.Errorf("[%s : tx.Rollback] : %s", functionName, rollbackErr)
			return customerror.GetError(customerror.InternalServer, rollbackErr)
		}
		return customerror.GetError(customerror.InternalServer, err)
	}

	if commitErr := tx.Commit(); commitErr != nil {
		c.l.Errorf("[%s : tx.Commit] : %s", functionName, commitErr)
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			c.l.Errorf("[%s : tx.Rollback] : %s", functionName, rollbackErr)
			return customerror.GetError(customerror.InternalServer, rollbackErr)
		}
		return customerror.GetError(customerror.InternalServer, commitErr)
	}

	return nil
}
