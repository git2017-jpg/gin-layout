package casbin

import (
	"context"
	"errors"
	"github.com/BooeZhang/gin-layout/internal/apiserver/model"
	"github.com/BooeZhang/gin-layout/internal/pkg/schema"
)

// UpdateCasBin 更新casBin 权限
func (csh *casBinService) UpdateCasBin(ctx context.Context, roleID string, casBinInfos []schema.CasBinInfoRes) error {
	csh.ClearCasBin(0, roleID)

	var rules [][]string
	for _, v := range casBinInfos {
		rules = append(rules, []string{roleID, v.Path, v.Method})
	}
	e := csh.CasBin()
	success, _ := e.AddPolicies(rules)
	if !success {
		return errors.New("存在相同api,添加失败,请联系管理员")
	}
	return nil
}

// UpdateCasBinApi 更新api的bashBin权限规则
func (csh *casBinService) UpdateCasBinApi(oldPath string, newPath string, oldMethod string, newMethod string) error {
	err := csh.store.GetDB().Table("casbin_rule").Model(&model.CasBinRuleModel{}).Where("v1 = ? AND v2 = ?", oldPath, oldMethod).Updates(map[string]interface{}{
		"v1": newPath,
		"v2": newMethod,
	}).Error
	return err
}
