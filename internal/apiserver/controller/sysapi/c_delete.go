package sysapi

import (
	"github.com/BooeZhang/gin-layout/internal/pkg/schema"
	"github.com/BooeZhang/gin-layout/pkg/erroron"
	"github.com/BooeZhang/gin-layout/pkg/log"
	"github.com/BooeZhang/gin-layout/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (u *Controller) Delete(c *gin.Context) {
	var idsInfo schema.ByIds
	err := c.ShouldBindJSON(&idsInfo)
	if err != nil {
		response.Ok(c, erroron.ErrParameter, nil)
		return
	}

	if err := u.srv.Delete(c.Request.Context(), idsInfo.Ids); err != nil {
		log.L(c).Error("删除api失败!", zap.Error(err))
		response.Ok(c, erroron.ErrParameter, nil)
	} else {
		response.Ok(c, nil, "删除成功")
	}
}
