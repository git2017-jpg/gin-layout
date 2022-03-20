package sysuser

import (
	"github.com/BooeZhang/gin-layout/internal/pkg/schema"
	"github.com/BooeZhang/gin-layout/pkg/erroron"
	"github.com/BooeZhang/gin-layout/pkg/request"
	"github.com/BooeZhang/gin-layout/pkg/response"
	"github.com/gin-gonic/gin"
)

func (u *Controller) SetUserRole(c *gin.Context) {
	var roleIds schema.ByIds
	err := c.ShouldBindJSON(&roleIds)
	if err != nil {
		response.Error(c, erroron.ErrParameter, nil)
	}
	userID := request.GetUserID(c)
	err = u.srv.SetUserRole(userID, roleIds.Ids)
	response.Ok(c, err, nil)
}
