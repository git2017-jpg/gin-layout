package sysuser

import (
	"context"
	"github.com/BooeZhang/gin-layout/internal/apiserver/model"
	"github.com/BooeZhang/gin-layout/pkg/log"
)

// Create 新建admin用户
func (u *sysUserService) Create(ctx context.Context, user *model.SysUserModel) error {
	var err error
	user.Password, err = user.Encrypt()
	if err != nil {
		log.L(ctx).Warn(err.Error())
	}
	err = u.store.GetDB().Create(user).Error
	if err != nil {
		return err
	}
	return nil
}
