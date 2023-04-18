package handler

import (
	"github.com/RyaWcksn/nann-e/api/v1/service"
	"github.com/RyaWcksn/nann-e/pkgs/logger"
)

type HandlerImpl struct {
	service service.IService
	log     logger.ILogger
}

type IHandler interface{}

func NewHandler(s service.IService, l logger.ILogger) *HandlerImpl {
	return &HandlerImpl{
		service: s,
		log:     l,
	}
}
