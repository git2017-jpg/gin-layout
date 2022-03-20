package sysrole

import (
	"github.com/BooeZhang/gin-layout/internal/apiserver/model"
	"github.com/BooeZhang/gin-layout/pkg/erroron"
	"github.com/BooeZhang/gin-layout/pkg/log"
	"github.com/BooeZhang/gin-layout/pkg/response"
	"github.com/gin-gonic/gin"
)

func (u *Controller) Update(c *gin.Context) {
	log.L(c).Info("user update function called.")

	var r model.RoleModel
	if err := c.ShouldBindJSON(&r); err != nil {
		response.Ok(c, erroron.ErrParameter, nil)
		return
	}

	err := u.srv.Update(c.Request.Context(), r)

	response.Ok(c, err, "创建成功")
}
