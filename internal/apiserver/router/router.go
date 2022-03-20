package router

import (
	docs "github.com/BooeZhang/gin-layout/docs"
	"github.com/BooeZhang/gin-layout/internal/pkg/options"
	"github.com/BooeZhang/gin-layout/middleware"
	"github.com/BooeZhang/gin-layout/pkg/erroron"
	"github.com/BooeZhang/gin-layout/pkg/log"
	"github.com/BooeZhang/gin-layout/pkg/response"
	"github.com/BooeZhang/gin-layout/store/mysql"
	"github.com/BooeZhang/gin-layout/store/redis"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter(g *gin.Engine) {
	installMiddleware(g)
	installController(g)
}

func installMiddleware(g *gin.Engine) {
}

func installController(g *gin.Engine) *gin.Engine {
	opts := options.GetOptions()
	storeIns, _ := mysql.GetMysqlFactoryOr(opts.MySQLOptions)
	storeCache, _ := redis.GetRedisFactoryOr(opts.RedisOptions)
	jwtStrategy := newJWTAuth(storeIns)
	errInit := jwtStrategy.MiddlewareInit()
	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}
	g.POST("/login", jwtStrategy.LoginHandler)
	g.POST("/logout", jwtStrategy.LogoutHandler)
	g.POST("/refresh", jwtStrategy.RefreshHandler)

	docs.SwaggerInfo.BasePath = "/v1"
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	g.NoRoute(jwtStrategy.MiddlewareFunc(), func(c *gin.Context) {
		response.Ok(c, erroron.ErrNotFound, nil)
	})
	{
		//indexController := index.NewIndexController(storeIns)
		//g.GET("/", indexController.Index)
		//g.GET("/login.html", indexController.Login)
		//g.GET("/welcome.html", jwtStrategy.MiddlewareFunc(), indexController.Welcome)
	}

	v1 := g.Group("/v1", jwtStrategy.MiddlewareFunc(), middleware.CasBinRbac())
	{
		installSysUserRouter(v1, storeIns, storeCache)
		installCasBinRouter(v1, storeIns, storeCache)
		installSysApiRouter(v1, storeIns, storeCache)
		installSysRoleRouter(v1, storeIns, storeCache)
	}

	return g
}
