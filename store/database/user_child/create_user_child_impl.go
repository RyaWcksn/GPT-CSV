package storechild

import (
	"context"
	dtochild "github.com/RyaWcksn/nann-e/dtos/user_child"
	customerror "github.com/RyaWcksn/nann-e/pkgs/error"
)

func (c *ChildImpl) CreateUserChild(ctx context.Context, payload *dtochild.CreateUserChildRequest) error {
	functionName := "ChildImpl.CreateUserChild"

	tx, err := c.DB.Begin()
	if err != nil {
		c.l.Errorf("[%s : c.DB.Begin - error begin transaction] : %v", functionName, err)
		return customerror.GetError(customerror.InternalServer, err)
	}

	stmt, err := tx.PrepareContext(ctx, QueryCreateUserChild)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			c.l.Errorf("update failed: %v, unable to back: %v", err, rollbackErr)
			return customerror.GetError(customerror.InternalServer, rollbackErr)
		}
		c.l.Errorf("[%s - tx.PrepareContext()] : %s", functionName, err)
		return customerror.GetError(customerror.InternalServer, err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		&payload.ParentId,
		&payload.ChildName,
		&payload.Age,
		&payload.RoleName,
	)

	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			c.l.Errorf("update failed: %v, unable to back: %v", err, rollbackErr)
			return customerror.GetError(customerror.InternalServer, rollbackErr)
		}

		c.l.Errorf("[%s - stmt.ExecContext]: %s", functionName, err)
		return customerror.GetError(customerror.InternalServer, err)
	}

	if commitErr := tx.Commit(); commitErr != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			c.l.Errorf("update failed: %v, unable to back: %v", err, rollbackErr)
			return customerror.GetError(customerror.InternalServer, rollbackErr)
		}
		c.l.Errorf("update failed: %v, unable to commit: %v", err, commitErr)
		return customerror.GetError(customerror.InternalServer, commitErr)
	}

	return nil
}
