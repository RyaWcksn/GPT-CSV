package storechild

import (
	"context"
	"encoding/json"
	"fmt"
	dtochild "github.com/RyaWcksn/nann-e/dtos/user_child"
	entitychild "github.com/RyaWcksn/nann-e/entities/user_child"
	customerror "github.com/RyaWcksn/nann-e/pkgs/error"
)

func (c *ChildImpl) GetOneUserChild(ctx context.Context, payload *dtochild.GetOneUserChildRequest) (*entitychild.UserChildDetail, error) {
	functionName := "ChildImpl.GetOneUserChild"
	res := new(entitychild.UserChildDetail)

	fmt.Println(payload)

	err := c.DB.QueryRowContext(ctx, QueryGetOneUserChild, payload.ParentId, payload.ChildName).Scan(
		&res.ChildName,
		&res.RoleName,
		&res.Age,
	)
	if err != nil {
		c.l.Errorf("[%s : u.DB.QueryRowContext]", functionName, err)
		return nil, customerror.GetError(customerror.InternalServer, err)
	}

	resByte, _ := json.Marshal(res)
	c.l.Debugf("Query Result : %s", resByte)

	return res, nil
}
