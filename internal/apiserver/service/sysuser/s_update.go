package sysuser

import (
	"context"
	"errors"
	"github.com/BooeZhang/gin-layout/internal/apiserver/model"
	"gorm.io/gorm"
)

func (u *sysUserService) Update(ctx context.Context, user model.SysUserModel) error {
	var (
		err error
		old model.SysUserModel
	)
	err = u.store.GetDB().Where("id = ?", user.ID).First(&old).Error
	if old.UserName != user.UserName {
		if !errors.Is(u.store.GetDB().Where("username = ? ", user.UserName).First(&model.SysUserModel{}).Error, gorm.ErrRecordNotFound) {
			return errors.New("存在相同用户名")
		}
	}
	if err != nil {
		return err
	} else {
		if user.Password != "" {
			user.Password, _ = user.Encrypt()
		}
		err = u.store.GetDB().Model(&model.SysUserModel{}).Where("id = ?", user.ID).Updates(&user).Error
	}

	return err
}
