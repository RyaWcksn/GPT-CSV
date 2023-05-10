package storechat

import (
	"database/sql"
	"github.com/RyaWcksn/nann-e/pkgs/logger"
)

type ChatImpl struct {
	DB *sql.DB
	l  logger.ILogger
}

func NewChatImpl(DB *sql.DB, l logger.ILogger) *ChatImpl {
	return &ChatImpl{
		DB: DB,
		l:  l,
	}
}
