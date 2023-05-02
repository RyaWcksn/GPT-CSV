package serviceroles

import (
	"context"
	"errors"
	dtoroles "github.com/RyaWcksn/nann-e/dtos/roles"
	entityroles "github.com/RyaWcksn/nann-e/entities/roles"
	customerror "github.com/RyaWcksn/nann-e/pkgs/error"
	"strings"
)

func (r *RolesService) CreateRoles(ctx context.Context, payload *dtoroles.CreateRoleRequest) (*entityroles.RoleDetails, error) {
	parentId := ctx.Value("ctxParentId").(string)
	payload.RoleId = payload.RoleName + "-" + parentId

	err := r.rolesRepo.CreateRoles(ctx, payload)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return nil, customerror.GetError(customerror.BadRequest, errors.New("role already exist, role name must be unique"))
		}
		return nil, err
	}

	res := entityroles.RoleDetails{
		Topic:            payload.Topic,
		Rules:            payload.Rules,
		Goals:            payload.Goals,
		ChildDescription: payload.ChildDescription,
		RoleName:         payload.RoleName,
		RoleDescription:  payload.RoleDescription,
	}

	return &res, nil
}
