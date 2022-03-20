package router

import (
	"github.com/BooeZhang/gin-layout/internal/apiserver/controller/sysrole"
	"github.com/BooeZhang/gin-layout/store"
	"github.com/gin-gonic/gin"
)

// 初始化sys_user相关路由
func installSysRoleRouter(g *gin.RouterGroup, db store.Factory, cache store.Cache) {
	role := g.Group("/sys-role")
	sysRoleController := sysrole.NewSysRoleController(db, cache)
	role.POST("add", sysRoleController.Create)
	role.POST("update", sysRoleController.Update)
	role.DELETE("delete", sysRoleController.Delete)
	role.GET("role-by-id", sysRoleController.GetSysRoleById)
	role.GET("role-list", sysRoleController.GetSysRoleList)
}
