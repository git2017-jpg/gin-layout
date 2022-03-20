package sysapi

import (
	"github.com/BooeZhang/gin-layout/internal/apiserver/service/casbin"
	"github.com/BooeZhang/gin-layout/internal/apiserver/service/sysapi"
	"github.com/BooeZhang/gin-layout/store"
)

// Controller 控制器
type Controller struct {
	srv sysapi.ISysApi
}

// NewSysApiController 系统管理员控制器
func NewSysApiController(db store.Factory, cache store.Cache) *Controller {
	casBin := casbin.NewCasBinService(db, cache)
	return &Controller{
		srv: sysapi.NewSysApiService(db, cache, casBin),
	}
}
