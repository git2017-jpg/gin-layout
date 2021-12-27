package sysuser

import (
	"github.com/BooeZhang/gin-layout/internal/apiserver/service"
	"github.com/BooeZhang/gin-layout/store"
)

type Controller struct {
	srv service.Service
}

func NewUserController(store store.Factory) *Controller {
	return &Controller{
		srv: service.NewService(store),
	}
}
