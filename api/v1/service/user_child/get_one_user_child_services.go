package servicechild

import (
	"context"
	dtochild "github.com/RyaWcksn/nann-e/dtos/user_child"
	entitychild "github.com/RyaWcksn/nann-e/entities/user_child"
)

func (c *ChildService) GetOneUserChild(ctx context.Context, payload *dtochild.GetOneUserChildRequest) (*entitychild.UserChildDetail, error) {
	parentId := ctx.Value("ctxParentId").(string)
	payload.ParentId = parentId
	
	childDetail, getChildErr := c.childRepo.GetOneUserChild(ctx, payload)
	if getChildErr != nil {
		return nil, getChildErr
	}

	return childDetail, nil
}

