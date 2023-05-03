package storechild

import (
	"database/sql"
	"github.com/RyaWcksn/nann-e/pkgs/logger"
)

type ChildImpl struct {
	DB *sql.DB
	l  logger.ILogger
}

func NewChildImpl(DB *sql.DB, l logger.ILogger) *ChildImpl {
	return &ChildImpl{
		DB: DB,
		l:  l,
	}
}
