package repochat

import (
	"context"
	dtochat "github.com/RyaWcksn/nann-e/dtos/chat"
	storechat "github.com/RyaWcksn/nann-e/store/database/chat"
)

type IRepository interface {
	CreateNewChat(ctx context.Context, payload *dtochat.CreateNewChatRequest) error
}

var _ IRepository = (*storechat.ChatImpl)(nil)
