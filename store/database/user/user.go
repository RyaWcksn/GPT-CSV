package user

import (
	"database/sql"

	"github.com/RyaWcksn/nann-e/pkgs/logger"
)

type UserImpl struct {
	DB  *sql.DB
	log logger.ILogger
}

func NewUser(db *sql.DB, log logger.ILogger) *UserImpl {
	return &UserImpl{
		DB:  db,
		log: log,
	}
}
