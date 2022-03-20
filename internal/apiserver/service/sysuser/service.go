package sysuser

import (
	"context"
	"github.com/BooeZhang/gin-layout/internal/apiserver/model"
	"github.com/BooeZhang/gin-layout/internal/pkg/schema"
	"github.com/BooeZhang/gin-layout/store"
)

type ISysUser interface {
	Create(ctx context.Context, user *model.SysUserModel) error
	Update(ctx context.Context, user model.SysUserModel) error
	Delete(ctx context.Context, ids []uint32) error
	GetSysUserById(ctx context.Context, id uint32) (user model.SysUserModel, err error)
	GetSysUserList(ctx context.Context, info schema.SearchSysUser) (list interface{}, total int64, err error)
	SetUserRole(userId uint32, roleId []uint32) (err error)
}

type sysUserService struct {
	store store.Factory
	cache store.Cache
}

var _ ISysUser = (*sysUserService)(nil)

// NewSysUserService 生成sys_user_service
func NewSysUserService(store store.Factory, cache store.Cache) *sysUserService {
	return &sysUserService{store: store, cache: cache}
}
