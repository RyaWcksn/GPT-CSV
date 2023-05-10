package handlerchat

import (
	servicechat "github.com/RyaWcksn/nann-e/api/v1/service/chat"
	"github.com/RyaWcksn/nann-e/pkgs/logger"
	"github.com/gofiber/fiber/v2"
)

type ChatHandler struct {
	chatService servicechat.IService
	l           logger.ILogger
}

func NewChatHandler(chatService servicechat.IService, l logger.ILogger) *ChatHandler {
	return &ChatHandler{
		chatService: chatService,
		l:           l,
	}
}

type IHandler interface {
	CreateNewChat(c *fiber.Ctx) error
}
