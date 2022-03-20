package sysrole

import (
	"context"
	"github.com/BooeZhang/gin-layout/internal/apiserver/model"
	"github.com/BooeZhang/gin-layout/internal/pkg/schema"
)

// GetSysRoleById 根据id获取角色
func (u *sysRoleService) GetSysRoleById(ctx context.Context, id uint32) (user model.RoleModel, err error) {
	err = u.store.GetDB().Where("id = ?", id).First(&user).Error
	return
}

func (u *sysRoleService) GetSysRoleList(ctx context.Context, info schema.SearchSysUser) (list interface{}, total int64, err error) {
	var apiList []model.RoleModel

	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := u.store.GetDB().Model(&model.RoleModel{})

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
