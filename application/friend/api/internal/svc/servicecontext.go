package svc

import (
	"im/application/friend/api/internal/config"
	"im/application/friend/rpc/friend"
	"im/application/user/rpc/user"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	UserRPC   user.User
	BizRedis  *redis.Redis
	FriendRPC friend.Friend
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UserRPC:   user.NewUser(zrpc.MustNewClient(c.UserRPC)),
		FriendRPC: friend.NewFriend(zrpc.MustNewClient(c.FriendRPC)),
		BizRedis:  redis.New(c.BizRedis.Host, redis.WithPass(c.BizRedis.Pass)),
	}
}
