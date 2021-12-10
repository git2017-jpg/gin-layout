package store

import (
	"gorm.io/gorm"
)

type Factory interface {
	GetMysql() *gorm.DB
	Close() error
}
