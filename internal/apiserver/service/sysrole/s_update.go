package sysrole

import (
	"context"
	"errors"
	"github.com/BooeZhang/gin-layout/internal/apiserver/model"
	"gorm.io/gorm"
)

func (u *sysRoleService) Update(ctx context.Context, role model.RoleModel) error {
	var (
		err error
		old model.RoleModel
	)
	err = u.store.GetDB().Where("id = ?", role.ID).First(&old).Error
	if old.Name != role.Name {
		if !errors.Is(u.store.GetDB().Where("name = ? ", role.Name).First(&model.RoleModel{}).Error, gorm.ErrRecordNotFound) {
			return errors.New("存在相同角色名")
		}
	}
	if err != nil {
		return err
	} else {
		err = u.store.GetDB().Model(&model.RoleModel{}).Where("id = ?", role.ID).Updates(&role).Error
	}

	return err
}
