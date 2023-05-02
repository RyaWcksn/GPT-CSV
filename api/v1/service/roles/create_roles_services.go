package serviceroles

import (
	"context"
	dtoroles "github.com/RyaWcksn/nann-e/dtos/roles"
	entityroles "github.com/RyaWcksn/nann-e/entities/roles"
)

func (r *RolesService) CreateRoles(ctx context.Context, payload *dtoroles.CreateRoleRequest) (*entityroles.CreateRoleDetails, error) {
	parentId := ctx.Value("ctxParentId").(string)
	payload.RoleId = payload.RoleName + "-" + parentId

	err := r.rolesRepo.CreateRoles(ctx, payload)
	if err != nil {
		return nil, err
	}

	res := entityroles.CreateRoleDetails{
		Topic:            payload.Topic,
		ChildDescription: payload.ChildDescription,
		RoleName:         payload.RoleName,
		RoleDescription:  payload.RoleDescription,
	}

	return &res, nil
}
