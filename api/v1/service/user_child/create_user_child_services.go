package servicechild

import (
	"context"
	"errors"
	dtochild "github.com/RyaWcksn/nann-e/dtos/user_child"
	entitychild "github.com/RyaWcksn/nann-e/entities/user_child"
	customerror "github.com/RyaWcksn/nann-e/pkgs/error"
	"strings"
)

func (c *ChildService) CreateUserChild(ctx context.Context, payload *dtochild.CreateUserChildRequest) (*entitychild.UserChildDetail, error) {
	parentId := ctx.Value("ctxParentId").(string)
	payload.ParentId = parentId

	err := c.childRepo.CreateUserChild(ctx, payload)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return nil, customerror.GetError(customerror.BadRequest, errors.New("child already exist, child name must be unique"))
		}
		return nil, err
	}

	res := entitychild.UserChildDetail{
		ChildName: payload.ChildName,
		RoleName:  payload.RoleName,
		Age:       payload.Age,
	}

	return &res, nil
}
