package service

import (
	"context"
	"github.com/BooeZhang/gin-layout/model"
	"github.com/BooeZhang/gin-layout/pkg/log"
	"github.com/BooeZhang/gin-layout/store"
)

type SysUser interface {
	Create(ctx context.Context, user *model.SysUserModel) error
}

type sysUserService struct {
	store store.Factory
}

var _ SysUser = (*sysUserService)(nil)

func newSysUser(srv *service) *sysUserService {
	return &sysUserService{store: srv.store}
}

func (u *sysUserService) Create(ctx context.Context, user *model.SysUserModel) error {
	var err error
	user.Password, err = user.Encrypt()
	if err != nil {
		log.L(ctx).Warn(err.Error())
	}
	err = u.store.GetMysql().Create(user).Error
	if err != nil {
		return err
	}
	return nil
}