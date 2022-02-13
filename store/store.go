package store

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type Factory interface {
	GetMysql() *gorm.DB
	Close() error
}

type Cache interface {
	GetCache() redis.UniversalClient
	Close() error
}
