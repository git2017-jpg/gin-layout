package redis

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/BooeZhang/gin-layout/internal/pkg/options"
	"github.com/BooeZhang/gin-layout/pkg/log"
	"github.com/BooeZhang/gin-layout/store"
	"os"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type datastore struct {
	redisCli redis.UniversalClient
}

func (ds *datastore) Close() error {
	if ds.redisCli != nil {
		return ds.redisCli.Close()
	}

	return nil
}

func (ds *datastore) GetRedis() redis.UniversalClient {
	return ds.redisCli
}

func (ds *datastore) GetMysql() *gorm.DB {
	return nil
}

var (
	redisFactory store.Factory
	once         sync.Once
)

type RedisOpts redis.UniversalOptions

func (o *RedisOpts) cluster() *redis.ClusterOptions {
	if len(o.Addrs) == 0 {
		o.Addrs = []string{"127.0.0.1:6379"}
	}

	return &redis.ClusterOptions{
		Addrs:     o.Addrs,
		OnConnect: o.OnConnect,

		Password: o.Password,

		MaxRedirects:   o.MaxRedirects,
		ReadOnly:       o.ReadOnly,
		RouteByLatency: o.RouteByLatency,
		RouteRandomly:  o.RouteRandomly,

		MaxRetries:      o.MaxRetries,
		MinRetryBackoff: o.MinRetryBackoff,
		MaxRetryBackoff: o.MaxRetryBackoff,

		DialTimeout:        o.DialTimeout,
		ReadTimeout:        o.ReadTimeout,
		WriteTimeout:       o.WriteTimeout,
		PoolSize:           o.PoolSize,
		MinIdleConns:       o.MinIdleConns,
		MaxConnAge:         o.MaxConnAge,
		PoolTimeout:        o.PoolTimeout,
		IdleTimeout:        o.IdleTimeout,
		IdleCheckFrequency: o.IdleCheckFrequency,

		TLSConfig: o.TLSConfig,
	}
}

func (o *RedisOpts) simple() *redis.Options {
	addr := "127.0.0.1:6379"
	if len(o.Addrs) > 0 {
		addr = o.Addrs[0]
	}

	return &redis.Options{
		Addr:      addr,
		OnConnect: o.OnConnect,

		DB:       o.DB,
		Password: o.Password,

		MaxRetries:      o.MaxRetries,
		MinRetryBackoff: o.MinRetryBackoff,
		MaxRetryBackoff: o.MaxRetryBackoff,

		DialTimeout:  o.DialTimeout,
		ReadTimeout:  o.ReadTimeout,
		WriteTimeout: o.WriteTimeout,

		PoolSize:           o.PoolSize,
		MinIdleConns:       o.MinIdleConns,
		MaxConnAge:         o.MaxConnAge,
		PoolTimeout:        o.PoolTimeout,
		IdleTimeout:        o.IdleTimeout,
		IdleCheckFrequency: o.IdleCheckFrequency,

		TLSConfig: o.TLSConfig,
	}
}

func (o *RedisOpts) failover() *redis.FailoverOptions {
	if len(o.Addrs) == 0 {
		o.Addrs = []string{"127.0.0.1:26379"}
	}

	return &redis.FailoverOptions{
		SentinelAddrs: o.Addrs,
		MasterName:    o.MasterName,
		OnConnect:     o.OnConnect,

		DB:       o.DB,
		Password: o.Password,

		MaxRetries:      o.MaxRetries,
		MinRetryBackoff: o.MinRetryBackoff,
		MaxRetryBackoff: o.MaxRetryBackoff,

		DialTimeout:  o.DialTimeout,
		ReadTimeout:  o.ReadTimeout,
		WriteTimeout: o.WriteTimeout,

		PoolSize:           o.PoolSize,
		MinIdleConns:       o.MinIdleConns,
		MaxConnAge:         o.MaxConnAge,
		PoolTimeout:        o.PoolTimeout,
		IdleTimeout:        o.IdleTimeout,
		IdleCheckFrequency: o.IdleCheckFrequency,

		TLSConfig: o.TLSConfig,
	}
}

// GetRedisFactoryOr 使用给定的配置创建 redis 工厂。
func GetRedisFactoryOr(opts *options.RedisOptions) (store.Factory, error) {
	if opts == nil && redisFactory == nil {
		return nil, fmt.Errorf("failed to get redis store fatory")
	}
	log.Debug("Creating new Redis connection pool")
	var (
		tlsConfig *tls.Config
		client    redis.UniversalClient
	)
	once.Do(func() {
		timeout := 5 * time.Second
		if opts.Timeout > 0 {
			timeout = time.Duration(opts.Timeout) * time.Second
		}
		// poolSize applies per cluster node and not for the whole cluster.
		poolSize := 500
		if opts.MaxActive > 0 {
			poolSize = opts.MaxActive
		}
		if opts.UseSSL {
			tlsConfig = &tls.Config{
				InsecureSkipVerify: opts.SSLInsecureSkipVerify,
			}
		}

		options := &RedisOpts{
			Addrs:        opts.Addrs,
			MasterName:   opts.MasterName,
			Password:     opts.Password,
			DB:           opts.Database,
			DialTimeout:  timeout,
			ReadTimeout:  timeout,
			WriteTimeout: timeout,
			IdleTimeout:  240 * timeout,
			PoolSize:     poolSize,
			TLSConfig:    tlsConfig,
		}

		if opts.MasterName != "" {
			log.Info("--> [REDIS] Creating sentinel-backed failover client")
			client = redis.NewFailoverClient(options.failover())
		} else if opts.EnableCluster {
			log.Info("--> [REDIS] Creating cluster client")
			client = redis.NewClusterClient(options.cluster())
		} else {
			log.Info("--> [REDIS] Creating single-node client")
			client = redis.NewClient(options.simple())
		}

		pong, err := client.Ping(context.Background()).Result()
		if err != nil {
			log.Error("redis connect ping failed, err:", zap.Any("err", err))
			os.Exit(1)
		} else {
			log.Info("redis connect ping response:", zap.String("pong", pong))
		}
		redisFactory = &datastore{client}
	})

	if redisFactory == nil {
		return nil, fmt.Errorf("failed to get redis store fatory, redisFactory: %+v", redisFactory)
	}

	return redisFactory, nil
}
