package casbin

import (
	"context"
	"github.com/BooeZhang/gin-layout/internal/pkg/options"
	"github.com/BooeZhang/gin-layout/internal/pkg/schema"
	"github.com/BooeZhang/gin-layout/store"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"sync"
)

type ICasBin interface {
	GetRoleApiPermissionsByRoleID(ctx context.Context, roleID uint32) (data []schema.CasBinInfoRes, err error)
	UpdateCasBin(ctx context.Context, roleID string, casBinInfos []schema.CasBinInfoRes) error
	UpdateCasBinApi(oldPath string, newPath string, oldMethod string, newMethod string) error
	ClearCasBin(v int, p ...string) bool
}

type casBinService struct {
	store store.Factory
	cache store.Cache
}

var _ ICasBin = (*casBinService)(nil)

var (
	syncedEnforcer *casbin.SyncedEnforcer
	once           sync.Once
)

func NewCasBinService(store store.Factory, cache store.Cache) *casBinService {
	return &casBinService{store: store, cache: cache}
}

// CasBin casBin配置
func (csh *casBinService) CasBin() *casbin.SyncedEnforcer {
	once.Do(func() {
		opt := options.GetOptions()
		a, _ := gormadapter.NewAdapterByDB(csh.store.GetDB())
		syncedEnforcer, _ = casbin.NewSyncedEnforcer(opt.CasBinOptions.ModelPath, a)
	})
	_ = syncedEnforcer.LoadPolicy()
	return syncedEnforcer
}
