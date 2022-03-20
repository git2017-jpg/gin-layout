package model

// CasBinRuleModel casbin 规则表
type CasBinRuleModel struct {
	PType  string `json:"p_type" gorm:"column:p_type"`
	RoleId string `json:"role_id" gorm:"column:v0"`
	Path   string `json:"path" gorm:"column:v1"`
	Method string `json:"method" gorm:"column:v2"`
}

func (CasBinRuleModel) TableName() string {
	return "casbin_rule"
}
