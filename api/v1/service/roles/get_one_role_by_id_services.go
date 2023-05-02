package serviceroles

import (
	"context"
	entityroles "github.com/RyaWcksn/nann-e/entities/roles"
)

func (r *RolesService) GetOneRoleById(ctx context.Context, roleName string) (*entityroles.RoleDetails, error) {
	parentId := ctx.Value("ctxParentId").(string)
	roleId := roleName + "-" + parentId

	roleDetail, getRoleErr := r.rolesRepo.GetOneRoleById(ctx, roleId)
	if getRoleErr != nil {
		return nil, getRoleErr
	}

	return roleDetail, nil
}
