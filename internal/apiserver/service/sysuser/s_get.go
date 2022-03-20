package sysuser

import (
	"context"
	"github.com/BooeZhang/gin-layout/internal/apiserver/model"
	"github.com/BooeZhang/gin-layout/internal/pkg/schema"
)

// GetSysUserById 根据id获取用户
func (u *sysUserService) GetSysUserById(ctx context.Context, id uint32) (user model.SysUserModel, err error) {
	err = u.store.GetDB().Where("id = ?", id).First(&user).Error
	return
}

func (u *sysUserService) GetSysUserList(ctx context.Context, info schema.SearchSysUser) (list interface{}, total int64, err error) {
	var apiList []model.SysUserModel

	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := u.store.GetDB().Model(&model.SysUserModel{})

	if info.Name != "" {
		db = db.Where("username LIKE ?", "%"+info.Name+"%")
	}

	err = db.Count(&total).Error

	if err != nil {
		return apiList, total, err
	}

	err = db.Limit(limit).Offset(offset).Find(&apiList).Error

	return apiList, total, err
}
