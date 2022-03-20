package sysrole

import (
	"github.com/BooeZhang/gin-layout/internal/apiserver/service/sysrole"
	"github.com/BooeZhang/gin-layout/store"
)

// Controller 控制器
type Controller struct {
	srv sysrole.ISysRole
}

// NewSysRoleController 系统角色控制器
func NewSysRoleController(store store.Factory, cache store.Cache) *Controller {
	return &Controller{
		srv: sysrole.NewSysRoleService(store, cache),
	}
}
