package sysuser

import (
	"github.com/BooeZhang/gin-layout/internal/apiserver/model"
	"gorm.io/gorm"
)

func (u *sysUserService) SetUserRole(userId uint32, roleId []uint32) (err error) {
	return u.store.GetDB().Transaction(func(tx *gorm.DB) error {
		TxErr := tx.Delete(&[]model.UserRoleModel{}, "user_id = ?", userId).Error
		if TxErr != nil {
			return TxErr
		}

		var userRole []model.UserRoleModel
		for _, v := range roleId {
			userRole = append(userRole, model.UserRoleModel{
				UserID: userId,
				RoleId: v,
			})
		}
		TxErr = tx.Create(&userRole).Error
		if TxErr != nil {
			return TxErr
		}
		// 返回 nil 提交事务
		return nil
	})
}
