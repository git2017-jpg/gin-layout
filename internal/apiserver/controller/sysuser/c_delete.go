package sysuser

import (
	"github.com/BooeZhang/gin-layout/internal/pkg/schema"
	"github.com/BooeZhang/gin-layout/pkg/erroron"
	"github.com/BooeZhang/gin-layout/pkg/response"
	"github.com/gin-gonic/gin"
)

func (u *Controller) Delete(c *gin.Context) {
	var idsInfo schema.ByIds
	err := c.ShouldBindJSON(&idsInfo)
	if err != nil {
		response.Ok(c, erroron.ErrParameter, nil)
		return
	}

	err = u.srv.Delete(c.Request.Context(), idsInfo.Ids)
	response.Ok(c, err, nil)
}
