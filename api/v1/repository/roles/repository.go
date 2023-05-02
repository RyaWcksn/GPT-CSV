package reporoles

import (
	"context"
	dtoroles "github.com/RyaWcksn/nann-e/dtos/roles"
	storeroles "github.com/RyaWcksn/nann-e/store/database/roles"
)

type IRepository interface {
	CreateRoles(ctx context.Context, payload *dtoroles.CreateRoleRequest) error
}

var _ IRepository = (*storeroles.RolesImpl)(nil)
