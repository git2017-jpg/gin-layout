package service

import "github.com/BooeZhang/gin-layout/store"

// Service 所有业务接口注册
type Service interface {
	SysUser() SysUser
	Index() Index
}

type service struct {
	store store.Factory
}

func NewService(store store.Factory) Service {
	return &service{store: store}
}

func (s *service) SysUser() SysUser {
	return newSysUser(s)
}

func (s *service) Index() Index {
	return newIndex(s)
}
