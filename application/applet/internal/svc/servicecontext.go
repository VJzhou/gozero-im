package svc

import (
	"im/application/applet/internal/config"
	"im/application/email/rpc/email"
	"im/application/user/rpc/user"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	UserRPC  user.User
	EmailRPC email.Email
	BizRedis *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		UserRPC:  user.NewUser(zrpc.MustNewClient(c.UserRPC)),
		EmailRPC: email.NewEmail(zrpc.MustNewClient(c.EmailRPC)),
		BizRedis: redis.New(c.BizRedis.Host, redis.WithPass(c.BizRedis.Pass)),
	}
}
