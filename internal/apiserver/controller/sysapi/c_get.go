package sysapi

import (
	"github.com/BooeZhang/gin-layout/internal/pkg/schema"
	"github.com/BooeZhang/gin-layout/pkg/erroron"
	"github.com/BooeZhang/gin-layout/pkg/response"
	"github.com/gin-gonic/gin"
)

func (u *Controller) GetApiById(c *gin.Context) {
	var idInfo schema.ById
	err := c.ShouldBindJSON(&idInfo)
	if err != nil {
		response.Ok(c, erroron.ErrParameter, nil)
		return
	}
	api, err := u.srv.GetApiById(c.Request.Context(), idInfo.ID)
	response.Ok(c, err, api)
}

func (u *Controller) GetAllApis(c *gin.Context) {
	apis, err := u.srv.GetAllApis(c.Request.Context())
	response.Ok(c, err, apis)
}

func (u *Controller) GetApiList(c *gin.Context) {
	var pageInfo schema.SearchApiParams
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.Ok(c, err, nil)
		return
	}

	list, total, err := u.srv.GetAPIInfoList(c.Request.Context(), pageInfo)
	response.PageOk(c, err, list, total, pageInfo.Page, pageInfo.PageSize)
}
