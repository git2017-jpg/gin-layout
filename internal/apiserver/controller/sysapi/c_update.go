package sysapi

import (
	"github.com/BooeZhang/gin-layout/internal/apiserver/model"
	"github.com/BooeZhang/gin-layout/pkg/erroron"
	"github.com/BooeZhang/gin-layout/pkg/log"
	"github.com/BooeZhang/gin-layout/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (u *Controller) Update(c *gin.Context) {
	var api model.SysApiModel
	err := c.ShouldBindJSON(&api)
	if err != nil {
		response.Ok(c, erroron.ErrParameter, nil)
		return
	}

	if err := u.srv.Update(c.Request.Context(), &api); err != nil {
		log.L(c).Error("更新api失败!", zap.Error(err))
		response.Ok(c, erroron.ErrParameter, nil)
	} else {
		response.Ok(c, nil, "更新api失败")
	}
}
