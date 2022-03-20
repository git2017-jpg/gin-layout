package sysrole

import (
	"context"
	"github.com/BooeZhang/gin-layout/internal/apiserver/model"
	"github.com/BooeZhang/gin-layout/pkg/log"
)

// Create 新建admin用户
func (u *sysRoleService) Create(ctx context.Context, role *model.RoleModel) error {
	var err error
	if err != nil {
		log.L(ctx).Warn(err.Error())
	}
	err = u.store.GetDB().Create(role).Error
	if err != nil {
		return err
	}
	return nil
}
