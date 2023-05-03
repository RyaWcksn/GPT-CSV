package servicechild

import (
	"context"
	dtochild "github.com/RyaWcksn/nann-e/dtos/user_child"
	entitychild "github.com/RyaWcksn/nann-e/entities/user_child"
)

func (c *ChildService) UpdateSingleUserChild(ctx context.Context, payload *dtochild.UpdateSingleUserChildRequest) (*entitychild.UserChildDetail, error) {
	parentId := ctx.Value("ctxParentId").(string)
	payload.ParentId = parentId

	err := c.childRepo.UpdateSingleUserChild(ctx, payload)
	if err != nil {
		return nil, err
	}

	res := entitychild.UserChildDetail{
		ChildName: payload.ChildName,
		Age:       payload.Age,
		RoleName:  payload.RoleName,
	}

	return &res, nil
}
