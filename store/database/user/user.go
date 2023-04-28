package storeusersparent

import (
	"database/sql"

	"github.com/RyaWcksn/nann-e/pkgs/logger"
)

type UserParentImpl struct {
	DB *sql.DB
	l  logger.ILogger
}

func NewUserParentImpl(DB *sql.DB, log logger.ILogger) *UserParentImpl {
	return &UserParentImpl{
		DB: DB,
		l:  log,
	}
}
