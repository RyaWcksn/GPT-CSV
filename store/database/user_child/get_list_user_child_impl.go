package storechild

import (
	"context"
	entitychild "github.com/RyaWcksn/nann-e/entities/user_child"
	customerror "github.com/RyaWcksn/nann-e/pkgs/error"
)

func (c *ChildImpl) GetListUserChild(ctx context.Context, offset, limit int) ([]entitychild.UserChildDetail, error) {
	functionName := "RolesImpl.GetListRole"

	children := make([]entitychild.UserChildDetail, 0)
	parentId := ctx.Value("ctxParentId").(string)

	rows, err := c.DB.QueryContext(ctx, QueryGetListChild, parentId, offset, limit)
	if err != nil {
		c.l.Errorf("[%s : r.DB.QueryContext] : %s", functionName, err)
		return nil, nil
	}

	for rows.Next() {
		var child entitychild.UserChildDetail
		if err = rows.Scan(
			&child.ChildName,
			&child.RoleName,
			&child.Age,
		); err != nil {
			c.l.Errorf("[%s : rows.Scan] : %s", functionName, err)
			return nil, customerror.GetError(customerror.InternalServer, err)
		}
		children = append(children, child)
	}

	return children, nil
}
