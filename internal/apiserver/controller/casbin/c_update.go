package casbin

import (
	"github.com/BooeZhang/gin-layout/internal/pkg/schema"
	"github.com/BooeZhang/gin-layout/pkg/log"
	"github.com/BooeZhang/gin-layout/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (u *Controller) UpdateCasBin(c *gin.Context) {
	var cmr schema.CasBinInReq
	_ = c.ShouldBindJSON(&cmr)

	if err := u.srv.UpdateCasBin(c.Request.Context(), cmr.RoleID, cmr.CasbinInfos); err != nil {
		log.Error("权限更新失败!", zap.Error(err))
		response.Error(c, err, nil)
		return
	} else {
		response.Ok(c, nil, "更新成功")
	}
}
