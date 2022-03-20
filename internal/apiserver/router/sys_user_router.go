package router

import (
	"github.com/BooeZhang/gin-layout/internal/apiserver/controller/sysuser"
	"github.com/BooeZhang/gin-layout/store"
	"github.com/gin-gonic/gin"
)

// 初始化sys_user相关路由
func installSysUserRouter(g *gin.RouterGroup, db store.Factory, cache store.Cache) {
	user := g.Group("/sys-user")
	sysUserController := sysuser.NewSysUserController(db, cache)
	user.POST("add", sysUserController.Create)
	user.POST("update", sysUserController.Update)
	user.DELETE("delete", sysUserController.Delete)
	user.GET("user-by-id", sysUserController.GetSysUserById)
	user.GET("user-list", sysUserController.GetSysUserList)
	user.POST("add-role", sysUserController.SetUserRole)
}
