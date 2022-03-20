package model

type SysApiModel struct {
	Model
	Path        string `json:"path" gorm:"comment:api路径"`
	Description string `json:"description" gorm:"comment:api中文描述"`
	ApiGroup    string `json:"api_group" gorm:"comment:api组"`
	Method      string `json:"method" gorm:"default:GET;comment:方法"`
}

func (SysApiModel) TableName() string {
	return "sys_api"
}
