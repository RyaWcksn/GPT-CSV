package serviceroles

import (
	"context"
	dtoroles "github.com/RyaWcksn/nann-e/dtos/roles"
	entityroles "github.com/RyaWcksn/nann-e/entities/roles"
)

func (r *RolesService) GetListRole(ctx context.Context, payload *dtoroles.GetListRoleRequest) ([]entityroles.RoleDetails, error) {
	offset := r.getOffset(payload)

	roleDetail, getListErr := r.rolesRepo.GetListRole(ctx, offset, payload.Limit)
	if getListErr != nil {
		return nil, getListErr
	}

	return roleDetail, nil
}

func (r *RolesService) getOffset(payload *dtoroles.GetListRoleRequest) int {
	var offset int
	pageNumber := payload.PageNumber - 1
	limit := payload.Limit

	if payload.PageNumber > 0 {
		offset = pageNumber * limit
	} else {
		offset = 0
	}

	return offset
}

