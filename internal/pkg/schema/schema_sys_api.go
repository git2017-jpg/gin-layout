package schema

// SearchApiParams api搜索参数
type SearchApiParams struct {
	PageInfo
	Path        string `json:"path" from:"Path"`
	Description string `json:"description" from:"description"`
	Method      string `json:"method" from:"method"`
	ApiGroup    string `json:"api_group" from:"api_group"`
}
