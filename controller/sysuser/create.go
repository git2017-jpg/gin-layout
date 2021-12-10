package sysuser

import (
	"github.com/BooeZhang/gin-layout/model"
	"github.com/BooeZhang/gin-layout/pkg/erroron"
	"github.com/BooeZhang/gin-layout/pkg/log"
	"github.com/BooeZhang/gin-layout/pkg/response"
	"github.com/gin-gonic/gin"
)

func (u *Controller) Create(c *gin.Context) {
	log.L(c).Info("user create function called.")

	var r model.SysUserModel
	if err := c.ShouldBindJSON(&r); err != nil {
		response.Ok(c, erroron.ErrParameter, nil)
		return
	}

	err := u.srv.SysUser().Create(c.Request.Context(), &r)
	if err != nil {
		response.Ok(c, erroron.ErrInternalServer, nil)
		return
	}

	response.Ok(c, nil, "创建成功")
}
