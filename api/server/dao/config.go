package dao

import (
	"mm-ndj/pkg/db/redis"

	clientv3 "go.etcd.io/etcd/client/v3"
	"gorm.io/gorm"
)

type CtxConfig struct {
	db   *gorm.DB
	rds  *redis.RedisServ
	etcd *clientv3.Client
}

type CtxOption func(conf *CtxConfig)

func NewServerCtx(options ...CtxOption) *ServiceCtx {
	c := &CtxConfig{}
	for _, option := range options {
		option(c)
	}
	return &ServiceCtx{Db: c.db}
}

func WithDB(db *gorm.DB) CtxOption {
	return func(conf *CtxConfig) {
		conf.db = db
	}
}
