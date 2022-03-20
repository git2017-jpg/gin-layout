package router

import (
	"github.com/BooeZhang/gin-layout/internal/apiserver/controller/sysapi"
	"github.com/BooeZhang/gin-layout/store"
	"github.com/gin-gonic/gin"
)

// 初始化sys_api相关路由
func installSysApiRouter(g *gin.RouterGroup, db store.Factory, cache store.Cache) {
	sysApiRouter := g.Group("/sys-api")
	sysApiService := sysapi.NewSysApiController(db, cache)

	sysApiRouter.POST("add", sysApiService.Create)
	sysApiRouter.DELETE("delete", sysApiService.Delete)
	sysApiRouter.POST("update", sysApiService.Update)
	sysApiRouter.GET("api-by-id", sysApiService.GetApiById) // 获取单条Api消息
	sysApiRouter.GET("all-api", sysApiService.GetAllApis)   // 获取所有api
	sysApiRouter.GET("api-list", sysApiService.GetApiList)  // 获取Api列表
}
