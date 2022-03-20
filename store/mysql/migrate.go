package mysql

import (
	"fmt"
	"github.com/BooeZhang/gin-layout/internal/apiserver/model"
	"gorm.io/gorm"
)

func migrateDatabase(db *gorm.DB) error {
	if err := db.AutoMigrate(
		new(model.SysUserModel),
		new(model.RoleModel),
		new(model.UserRoleModel),
		new(model.CasBinRuleModel),
		new(model.SysApiModel),
	); err != nil {
		return fmt.Errorf("migrate user model failed: %w", err)
	}
	return nil
}
