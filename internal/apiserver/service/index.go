package service

import (
	"context"
	"github.com/BooeZhang/gin-layout/store"
)

type Index interface {
	Index(ctx context.Context) error
}

type indexService struct {
	store store.Factory
}

var _ Index = (*indexService)(nil)

func newIndex(srv *service) *indexService {
	return &indexService{store: srv.store}
}

func (i *indexService) Index(ctx context.Context) error {
	return nil
}
