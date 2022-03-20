package schema

// CasBinInfoRes casBin 响应
type CasBinInfoRes struct {
	Path   string `json:"path"`   // 路径
	Method string `json:"method"` // 方法
}

// CasBinInReq casBin 请求
type CasBinInReq struct {
	RoleID      string          `json:"role_id"` // 权限id
	CasbinInfos []CasBinInfoRes `json:"casbinInfos"`
}
