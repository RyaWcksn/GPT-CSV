package reporoles

import (
	"context"
	dtoroles "github.com/RyaWcksn/nann-e/dtos/roles"
	entityroles "github.com/RyaWcksn/nann-e/entities/roles"
	storeroles "github.com/RyaWcksn/nann-e/store/database/roles"
)

type IRepository interface {
	CreateRoles(ctx context.Context, payload *dtoroles.CreateRoleRequest) error
	GetOneRole(ctx context.Context, payload *dtoroles.GetOneRoleRequest) (*entityroles.RoleDetails, error)
	GetListRole(ctx context.Context, offset, limit int) ([]entityroles.RoleDetails, error)
	UpdateSingleRoleById(ctx context.Context, payload *dtoroles.UpdateSingleRoleRequest) error
}

var _ IRepository = (*storeroles.RolesImpl)(nil)
