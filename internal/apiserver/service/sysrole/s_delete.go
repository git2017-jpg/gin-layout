package sysrole

import (
	"context"
	"github.com/BooeZhang/gin-layout/internal/apiserver/model"
	"github.com/BooeZhang/gin-layout/pkg/log"
)

// Delete 删除用户
func (u *sysRoleService) Delete(ctx context.Context, ids []uint32) error {
	var (
		err error
	)
	err = u.store.GetDB().Delete(&[]model.RoleModel{}, "id in ?", ids).Error
	if err != nil {
		log.L(ctx).Warn(err.Error())
		return err
	}
	return nil
}
