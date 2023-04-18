package service

import (
	"github.com/RyaWcksn/nann-e/api/v1/repository"
	"github.com/RyaWcksn/nann-e/pkgs/logger"
)

type ServiceImpl struct {
	user   repository.IUser
	prompt repository.IPrompt
	openAi repository.IOpenAi
	log    logger.ILogger
}

type IService interface{}

func NewService(u repository.IUser, p repository.IPrompt, a repository.IOpenAi, l logger.ILogger) *ServiceImpl {
	return &ServiceImpl{
		user:   u,
		prompt: p,
		openAi: a,
		log:    l,
	}
}
