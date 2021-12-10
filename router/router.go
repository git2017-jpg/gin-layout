package router

import (
	"fmt"
	"github.com/BooeZhang/gin-layout/internal/apiserver/options"
	"github.com/BooeZhang/gin-layout/pkg/erroron"
	"github.com/BooeZhang/gin-layout/pkg/log"
	"github.com/BooeZhang/gin-layout/pkg/response"
	"github.com/BooeZhang/gin-layout/store/mysql"
	"github.com/gin-gonic/gin"
)

func InitRouter(g *gin.Engine) {
	installMiddleware(g)
	installController(g)
}

func installMiddleware(g *gin.Engine) {
}

func installController(g *gin.Engine) *gin.Engine {
	opts := options.GetOptions()
	fmt.Println(opts.MySQLOptions)
	storeIns, _ := mysql.GetMysqlFactoryOr(opts.MySQLOptions)
	jwtStrategy := newJWTAuth(storeIns)
	errInit := jwtStrategy.MiddlewareInit()
	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}
	g.POST("/login", jwtStrategy.LoginHandler)
	g.POST("/logout", jwtStrategy.LogoutHandler)
	g.POST("/refresh", jwtStrategy.RefreshHandler)

	g.NoRoute(jwtStrategy.MiddlewareFunc(), func(c *gin.Context) {
		response.Ok(c, erroron.ErrNotFound, nil)
	})
	{
		//indexController := index.NewIndexController(storeIns)
		//g.GET("/", indexController.Index)
		//g.GET("/login.html", indexController.Login)
		//g.GET("/welcome.html", jwtStrategy.MiddlewareFunc(), indexController.Welcome)
	}

	//v1 := g.Group("/v1", jwtStrategy.MiddlewareFunc())
	//{
	//	user := v1.Group("/user")
	//	userController := sysuser.NewUserController(storeIns)
	//	user.POST("add", userController.Create)
	//}

	return g
}