package serviceroles

import (
	"context"
	dtoroles "github.com/RyaWcksn/nann-e/dtos/roles"
	entityroles "github.com/RyaWcksn/nann-e/entities/roles"
)

func (r *RolesService) UpdateSingleRoleById(ctx context.Context, payload *dtoroles.UpdateSingleRoleRequest) (*entityroles.RoleDetails, error) {
	parentId := ctx.Value("ctxParentId").(string)
	payload.ParentId = parentId

	err := r.rolesRepo.UpdateSingleRoleById(ctx, payload)
	if err != nil {
		return nil, err
	}

	res := entityroles.RoleDetails{
		RoleName:         payload.RoleName,
		Topic:            payload.Topic,
		Rules:            payload.Rules,
		Goals:            payload.Goals,
		ChildDescription: payload.ChildDescription,
		RoleDescription:  payload.RoleDescription,
	}

	return &res, nil
}
