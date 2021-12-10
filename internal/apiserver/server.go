package apiserver

import (
	"github.com/BooeZhang/gin-layout/internal/apiserver/config"
	options2 "github.com/BooeZhang/gin-layout/internal/pkg/options"
	"github.com/BooeZhang/gin-layout/internal/pkg/server"
	"github.com/BooeZhang/gin-layout/router"
)

type apiServer struct {
	redisOptions *options2.RedisOptions
	mysqlOptions *options2.MySQLOptions
	// etcdOptions      *genericoptions.EtcdOptions
	gRPCAPIServer    *grpcAPIServer
	apiServer *server.APIServer
}

func buildGenericConfig(cfg *config.Config) (genericConfig *server.Config, lastErr error) {
	genericConfig = server.NewConfig()
	// 服务器运行时配置
	if lastErr = cfg.ServerRunOptions.ApplyTo(genericConfig); lastErr != nil {
		return
	}

	if lastErr = cfg.FeatureOptions.ApplyTo(genericConfig); lastErr != nil {
		return
	}

	if lastErr = cfg.HttpServerOptions.ApplyTo(genericConfig); lastErr != nil {
		return
	}

	return
}

// createAPIServer 创建api服务器
func createAPIServer(cfg *config.Config) (*apiServer, error) {
	genericConfig, err := buildGenericConfig(cfg)
	if err != nil {
		return nil, err
	}

	genericServer, err := genericConfig.Complete().New()
	if err != nil {
		return nil, err
	}

	return &apiServer{
		redisOptions:     cfg.RedisOptions,
		apiServer: genericServer,
		mysqlOptions:     cfg.MySQLOptions,
		// etcdOptions:      cfg.EtcdOptions,
	}, nil
}

func (s *apiServer) Run() error {
	router.InitRouter(s.apiServer.Engine)
	return s.apiServer.Run()
}
