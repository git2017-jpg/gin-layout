package model

// RoleModel 角色表
type RoleModel struct {
	Model
	Name   string `json:"name" gorm:"not null;unique;comment:角色名"`
	Remark string `json:"remark" gorm:"comment:备注"`
}

func (RoleModel) TableName() string {
	return "role"
}

// UserRoleModel 用户与角色绑定关系表
type UserRoleModel struct {
	Model
	UserID uint32 `json:"user_id" gorm:"not null;comment:用户id"`
	RoleId uint32 `json:"role_id" gorm:"not null;comment:角色id"`
}

func (UserRoleModel) TableName() string {
	return "user_role"
}
