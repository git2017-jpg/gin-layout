package schema

// LoginInfoRes 登录
type LoginInfoRes struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
