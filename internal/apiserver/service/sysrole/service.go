package sysrole

import (
	"context"
	"github.com/BooeZhang/gin-layout/internal/apiserver/model"
	"github.com/BooeZhang/gin-layout/internal/pkg/schema"
	"github.com/BooeZhang/gin-layout/store"
)

type ISysRole interface {
	Create(ctx context.Context, user *model.RoleModel) error
	Update(ctx context.Context, user model.RoleModel) error
	Delete(ctx context.Context, ids []uint32) error
	GetSysRoleById(ctx context.Context, id uint32) (user model.RoleModel, err error)
	GetSysRoleList(ctx context.Context, info schema.SearchSysUser) (list interface{}, total int64, err error)
}

type sysRoleService struct {
	store store.Factory
	cache store.Cache
}

var _ ISysRole = (*sysRoleService)(nil)

// NewSysRoleService 生成sys_role_service
func NewSysRoleService(store store.Factory, cache store.Cache) *sysRoleService {
	return &sysRoleService{store: store, cache: cache}
}
