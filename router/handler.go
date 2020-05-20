package router

import (
	"github.com/sukstar76/go-kafka/service"
)

type Handler struct {
	us service.UserServiceInterface
	ls service.LogServiceInterface
}

func NewHandler(us service.UserServiceInterface, ls service.LogServiceInterface) *Handler {
	return &Handler{
		us: us,
		ls: ls,
	}
}
