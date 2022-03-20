package sysapi

import (
	"context"
	"github.com/BooeZhang/gin-layout/internal/apiserver/model"
	"github.com/BooeZhang/gin-layout/pkg/log"
)

// Create 新建api
func (u *sysApiService) Create(ctx context.Context, api *model.SysApiModel) error {
	var err error
	if err != nil {
		log.L(ctx).Warn(err.Error())
	}
	err = u.store.GetDB().Create(api).Error
	if err != nil {
		return err
	}
	return nil
}
