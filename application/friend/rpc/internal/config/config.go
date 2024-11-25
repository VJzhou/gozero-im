package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

type Config struct {
	zrpc.RpcServerConf
	Consul consul.Conf
	DB     struct {
		DataSource   string
		MaxOpenConns int `json:,default=10`
		MaxIdleConns int `json:,default=100`
		MaxLifeTime  int `json:,default=3600`
	}
	BizRedis   redis.RedisConf
	CacheRedis cache.CacheConf
}
