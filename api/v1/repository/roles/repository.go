package reporoles

import (
	"context"
	dtoroles "github.com/RyaWcksn/nann-e/dtos/roles"
	entityroles "github.com/RyaWcksn/nann-e/entities/roles"
	storeroles "github.com/RyaWcksn/nann-e/store/database/roles"
)

type IRepository interface {
	CreateRoles(ctx context.Context, payload *dtoroles.CreateRoleRequest) error
	GetOneRoleById(ctx context.Context, roleId string) (*entityroles.RoleDetails, error)
	GetListRole(ctx context.Context, offset, limit int) ([]entityroles.RoleDetails, error)
}

var _ IRepository = (*storeroles.RolesImpl)(nil)
