package prompt

import (
	"database/sql"

	"github.com/RyaWcksn/nann-e/pkgs/logger"
)

type PromptImpl struct {
	DB  *sql.DB
	log logger.ILogger
}

func NewPrompt(db *sql.DB, log logger.ILogger) *PromptImpl {
	return &PromptImpl{
		DB:  db,
		log: log,
	}
}
