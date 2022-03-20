package sysapi

import (
	"context"
	"github.com/BooeZhang/gin-layout/internal/apiserver/model"
	"github.com/BooeZhang/gin-layout/pkg/log"
)

// Delete 删除api
func (u *sysApiService) Delete(ctx context.Context, ids []uint32) error {
	var (
		apis []model.SysApiModel
		err  error
	)
	err = u.store.GetDB().Find(&apis, ids).Error
	err = u.store.GetDB().Delete(&[]model.SysApiModel{}, "id in ?", ids).Error
	if err != nil {
		log.L(ctx).Warn(err.Error())
		return err
	}
	for _, api := range apis {
		u.casBin.ClearCasBin(1, api.Path, api.Method)
	}
	return nil
}
