package servicechat

import (
	"context"
	repochat "github.com/RyaWcksn/nann-e/api/v1/repository/chat"
	reporoles "github.com/RyaWcksn/nann-e/api/v1/repository/roles"
	repochild "github.com/RyaWcksn/nann-e/api/v1/repository/user_child"
	dtochat "github.com/RyaWcksn/nann-e/dtos/chat"
	entitychat "github.com/RyaWcksn/nann-e/entities/chat"
	"github.com/RyaWcksn/nann-e/pkgs/logger"
)

type ChatService struct {
	chatRepo  repochat.IRepository
	childRepo repochild.IRepository
	rolesRepo reporoles.IRepository
	l         logger.ILogger
}

func NewChatService(chatRepo repochat.IRepository, childRepo repochild.IRepository, rolesRepo reporoles.IRepository, l logger.ILogger) *ChatService {
	return &ChatService{
		chatRepo:  chatRepo,
		childRepo: childRepo,
		rolesRepo: rolesRepo,
		l:         l,
	}
}

type IService interface {
	CreateNewChat(ctx context.Context, payload *dtochat.CreateNewChatRequest) (*entitychat.CreateNewChatDetail, error)
}

var _ IService = (*ChatService)(nil)
