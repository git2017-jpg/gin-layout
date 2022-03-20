package sysapi

import (
	"context"
	"github.com/BooeZhang/gin-layout/internal/apiserver/model"
	"github.com/BooeZhang/gin-layout/internal/apiserver/service/casbin"
	"github.com/BooeZhang/gin-layout/internal/pkg/schema"
	"github.com/BooeZhang/gin-layout/store"
)

type ISysApi interface {
	Create(ctx context.Context, api *model.SysApiModel) error
	Delete(ctx context.Context, ids []uint32) error
	Update(ctx context.Context, api *model.SysApiModel) error
	GetApiById(ctx context.Context, id uint32) (api model.SysApiModel, err error)
	GetAllApis(ctx context.Context) (apis []model.SysApiModel, err error)
	GetAPIInfoList(ctx context.Context, info schema.SearchApiParams) (list interface{}, total int64, err error)
}

type sysApiService struct {
	store  store.Factory
	cache  store.Cache
	casBin casbin.ICasBin
}

var _ ISysApi = (*sysApiService)(nil)

// NewSysApiService 生成sys_api_service
func NewSysApiService(store store.Factory, cache store.Cache, casBin casbin.ICasBin) *sysApiService {
	return &sysApiService{store: store, cache: cache, casBin: casBin}
}
