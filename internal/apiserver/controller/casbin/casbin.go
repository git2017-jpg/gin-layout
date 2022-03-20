package casbin

import (
	"github.com/BooeZhang/gin-layout/internal/apiserver/service/casbin"
	"github.com/BooeZhang/gin-layout/store"
)

type Controller struct {
	srv casbin.ICasBin
}

// NewCasBinController 初始化casBin控制器
func NewCasBinController(store store.Factory, cache store.Cache) *Controller {
	return &Controller{
		srv: casbin.NewCasBinService(store, cache),
	}
}
