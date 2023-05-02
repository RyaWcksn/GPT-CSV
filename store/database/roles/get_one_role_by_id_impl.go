package storeroles

import (
	"context"
	"encoding/json"
	entityroles "github.com/RyaWcksn/nann-e/entities/roles"
	customerror "github.com/RyaWcksn/nann-e/pkgs/error"
)

func (r *RolesImpl) GetOneRoleById(ctx context.Context, roleId string) (*entityroles.RoleDetails, error) {
	functionName := "RolesImpl.GetOneRoleById"
	res := new(entityroles.RoleDetails)

	err := r.DB.QueryRowContext(ctx, QueryGetOneRoleById, roleId).Scan(
		&res.Topic,
		&res.Rules,
		&res.Goals,
		&res.ChildDescription,
		&res.RoleName,
		&res.RoleDescription,
	)
	if err != nil {
		r.l.Errorf("[%s : u.DB.QueryRowContext]", functionName, err)
		return nil, customerror.GetError(customerror.InternalServer, err)
	}

	resByte, _ := json.Marshal(res)
	r.l.Debugf("Query Result : %s", resByte)

	return res, nil
}
