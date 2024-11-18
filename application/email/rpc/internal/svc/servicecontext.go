package svc

import (
	"im/application/email/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config   config.Config
	BizRedis *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {

	rds := redis.MustNewRedis(redis.RedisConf{
		Host: c.BizRedis.Host,
		Pass: c.BizRedis.Pass,
		Type: c.BizRedis.Type,
	})

	return &ServiceContext{
		Config:   c,
		BizRedis: rds,
	}
}
