package schema

type SearchSysUser struct {
	PageInfo
	Name string `json:"name" from:"name"`
}
