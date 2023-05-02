package storeroles

import (
	"database/sql"
	"github.com/RyaWcksn/nann-e/pkgs/logger"
)

type RolesImpl struct {
	DB *sql.DB
	l  logger.ILogger
}

func NewRolesImpl(DB *sql.DB, l logger.ILogger) *RolesImpl {
	return &RolesImpl{
		DB: DB,
		l:  l,
	}
}
