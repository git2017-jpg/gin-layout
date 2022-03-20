package casbin

import (
	"context"
	"github.com/BooeZhang/gin-layout/internal/pkg/schema"
	"strconv"
)

// GetRoleApiPermissionsByRoleID 根据角色id获取api权限
func (csh *casBinService) GetRoleApiPermissionsByRoleID(ctx context.Context, roleID uint32) (data []schema.CasBinInfoRes, err error) {
	e := csh.CasBin()
	list := e.GetFilteredPolicy(0, strconv.FormatUint(uint64(roleID), 10))
	for _, v := range list {
		data = append(data, schema.CasBinInfoRes{
			Path:   v[1],
			Method: v[2],
		})
	}
	return data, nil
}
