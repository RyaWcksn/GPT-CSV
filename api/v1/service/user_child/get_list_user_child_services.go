package servicechild

import (
	"context"
	dtochild "github.com/RyaWcksn/nann-e/dtos/user_child"
	entitychild "github.com/RyaWcksn/nann-e/entities/user_child"
)

func (c *ChildService) GetListUserChild(ctx context.Context, payload *dtochild.GetListUserChildRequest) ([]entitychild.UserChildDetail, error) {
	offset := c.getOffset(payload)

	childDetail, getListErr := c.childRepo.GetListUserChild(ctx, offset, payload.Limit)
	if getListErr != nil {
		return nil, getListErr
	}

	return childDetail, nil
}

func (c *ChildService) getOffset(payload *dtochild.GetListUserChildRequest) int {
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
