package sysapi

import (
	"context"
	"errors"
	"github.com/BooeZhang/gin-layout/internal/apiserver/model"
	"gorm.io/gorm"
)

// Update 更新api
func (u *sysApiService) Update(ctx context.Context, api *model.SysApiModel) error {
	var (
		err error
		old model.SysApiModel
	)
	err = u.store.GetDB().Where("id = ?", api.ID).First(&old).Error
	if old.Path != api.Path || old.Method != api.Method {
		if !errors.Is(u.store.GetDB().Where("path = ? AND method = ?", api.Path, api.Method).First(&model.SysApiModel{}).Error, gorm.ErrRecordNotFound) {
			return errors.New("存在相同api路径")
		}
	}
	if err != nil {
		return err
	} else {
		err = u.casBin.UpdateCasBinApi(old.Path, api.Path, old.Method, api.Method)
		if err != nil {
			return err
		} else {
			err = u.store.GetDB().Save(&api).Error
		}
	}
	return err
}
