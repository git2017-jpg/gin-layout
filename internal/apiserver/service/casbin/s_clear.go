package casbin

// ClearCasBin 清除匹配的权限
func (csh *casBinService) ClearCasBin(v int, p ...string) bool {
	e := csh.CasBin()
	success, _ := e.RemoveFilteredPolicy(v, p...)
	return success
}
