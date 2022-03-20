package router

import (
	"github.com/BooeZhang/gin-layout/internal/apiserver/controller/casbin"
	"github.com/BooeZhang/gin-layout/store"
	"github.com/gin-gonic/gin"
)

// 初始化casBin相关路由
func installCasBinRouter(g *gin.RouterGroup, db store.Factory, cache store.Cache) {
	casBinRouter := g.Group("/casbin")
	casBinService := casbin.NewCasBinController(db, cache)

	casBinRouter.GET("get-api-by-role-id", casBinService.GetRoleApiPermissionsByRoleID)
	casBinRouter.POST("update-casbin", casBinService.UpdateCasBin)
}
