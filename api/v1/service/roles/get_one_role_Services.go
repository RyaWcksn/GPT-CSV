package serviceroles

import (
	"context"
	dtoroles "github.com/RyaWcksn/nann-e/dtos/roles"
	entityroles "github.com/RyaWcksn/nann-e/entities/roles"
)

func (r *RolesService) GetOneRole(ctx context.Context, payload *dtoroles.GetOneRoleRequest) (*entityroles.RoleDetails, error) {
	parentId := ctx.Value("ctxParentId").(string)
	payload.ParentId = parentId

	roleDetail, getRoleErr := r.rolesRepo.GetOneRole(ctx, payload)
	if getRoleErr != nil {
		return nil, getRoleErr
	}

	return roleDetail, nil
}
