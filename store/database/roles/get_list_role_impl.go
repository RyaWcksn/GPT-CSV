package storeroles

import (
	"context"
	entityroles "github.com/RyaWcksn/nann-e/entities/roles"
	customerror "github.com/RyaWcksn/nann-e/pkgs/error"
)

func (r *RolesImpl) GetListRole(ctx context.Context, offset, limit int) ([]entityroles.RoleDetails, error) {
	functionName := "RolesImpl.GetListRole"

	roles := make([]entityroles.RoleDetails, 0)
	parentId := ctx.Value("ctxParentId").(string)

	rows, err := r.DB.QueryContext(ctx, QueryGetListRole, parentId, offset, limit)
	if err != nil {
		r.l.Errorf("[%s : r.DB.QueryContext] : %s", functionName, err)
		return nil, nil
	}

	for rows.Next() {
		var role entityroles.RoleDetails
		if err = rows.Scan(
			&role.RoleName,
			&role.Topic,
			&role.Rules,
			&role.Goals,
			&role.ChildDescription,
			&role.RoleDescription,
		); err != nil {
			r.l.Errorf("[%s : rows.Scan] : %s", functionName, err)
			return nil, customerror.GetError(customerror.InternalServer, err)
		}
		roles = append(roles, role)
	}

	return roles, nil
}
