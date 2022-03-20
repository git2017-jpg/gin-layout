package casbin

import (
	"github.com/BooeZhang/gin-layout/pkg/log"
	"github.com/BooeZhang/gin-layout/pkg/request"
	"github.com/BooeZhang/gin-layout/pkg/response"
	"github.com/gin-gonic/gin"
)

func (u *Controller) GetRoleApiPermissionsByRoleID(c *gin.Context) {
	log.L(c).Info("get role api permissions by role_id function called.")

	userId := request.GetUserID(c)
	data, err := u.srv.GetRoleApiPermissionsByRoleID(c.Request.Context(), userId)
	if err != nil {
		response.Error(c, err, nil)
		return
	}

	response.Ok(c, nil, data)
}
