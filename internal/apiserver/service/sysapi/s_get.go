package sysapi

import (
	"context"
	"github.com/BooeZhang/gin-layout/internal/apiserver/model"
	"github.com/BooeZhang/gin-layout/internal/pkg/schema"
)

// GetApiById 根据id获取api
func (u *sysApiService) GetApiById(ctx context.Context, id uint32) (api model.SysApiModel, err error) {
	err = u.store.GetDB().Where("id = ?", id).First(&api).Error
	return
}

// GetAllApis 获取全部api
func (u *sysApiService) GetAllApis(ctx context.Context) (apis []model.SysApiModel, err error) {
	err = u.store.GetDB().Find(&apis).Error
	return
}

func (u *sysApiService) GetAPIInfoList(ctx context.Context, info schema.SearchApiParams) (list interface{}, total int64, err error) {
	var apiList []model.SysApiModel

	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := u.store.GetDB().Model(&model.SysApiModel{})

	if info.Path != "" {
		db = db.Where("path LIKE ?", "%"+info.Path+"%")
	}

	if info.Description != "" {
		db = db.Where("description LIKE ?", "%"+info.Description+"%")
	}

	if info.Method != "" {
		db = db.Where("method = ?", info.Method)
	}

	if info.ApiGroup != "" {
		db = db.Where("api_group = ?", info.ApiGroup)
	}

	err = db.Count(&total).Error

	if err != nil {
		return apiList, total, err
	}

	err = db.Limit(limit).Offset(offset).Find(&apiList).Error

	return apiList, total, err
}
