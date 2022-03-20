package sysuser

import (
	"github.com/BooeZhang/gin-layout/internal/apiserver/service/sysuser"
	"github.com/BooeZhang/gin-layout/store"
)

// Controller 控制器
type Controller struct {
	srv sysuser.ISysUser
}

// NewSysUserController 系统管理员控制器
func NewSysUserController(store store.Factory, cache store.Cache) *Controller {
	return &Controller{
		srv: sysuser.NewSysUserService(store, cache),
	}
}
