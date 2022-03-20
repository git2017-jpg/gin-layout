package middleware

import (
	"github.com/BooeZhang/gin-layout/internal/apiserver/service/casbin"
	"github.com/BooeZhang/gin-layout/pkg/erroron"
	"github.com/BooeZhang/gin-layout/pkg/response"
	"github.com/BooeZhang/gin-layout/store/mysql"
	"github.com/BooeZhang/gin-layout/store/redis"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var (
	db    = mysql.GetMysqlFactory()
	cache = redis.GetRedisFactory()
)

var casbinService = casbin.NewCasBinService(db, cache)

// CasBinRbac 拦截器
func CasBinRbac() gin.HandlerFunc {
	return func(c *gin.Context) {
		waitUse := jwt.ExtractClaims(c)
		// 获取请求的PATH
		obj := c.Request.URL.Path
		// 获取请求方法
		act := c.Request.Method
		// 获取用户的角色
		sub, _ := waitUse["role_ids"].([]uint32)
		e := casbinService.CasBin()
		// 判断策略中是否存在
		for _, role := range sub {
			success, _ := e.Enforce(role, obj, act)
			if success {
				c.Next()
			} else {
				response.Ok(c, erroron.ErrNoPerm, nil)
				c.Abort()
				return
			}
		}
	}
}
