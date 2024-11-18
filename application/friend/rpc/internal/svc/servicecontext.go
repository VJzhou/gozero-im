package svc

import (
	"im/application/friend/rpc/internal/config"
	"im/application/friend/rpc/internal/model"
	"im/pkg/orm"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config           config.Config
	DB               *orm.DB
	BizRedis         *redis.Redis
	ApplyFriendModel *model.ApplyFriendModel
	FriendModel      *model.FriendModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	db := orm.MustNewMySQL(&orm.Config{
		DSN:          c.DB.DataSource,
		MaxOpenConns: c.DB.MaxOpenConns,
		MaxIdleConns: c.DB.MaxIdleConns,
		MaxLifetime:  c.DB.MaxLifeTime,
	})

	rds := redis.MustNewRedis(redis.RedisConf{
		Host: c.BizRedis.Host,
		Pass: c.BizRedis.Pass,
		Type: c.BizRedis.Type,
	})

	return &ServiceContext{
		Config:           c,
		DB:               db,
		BizRedis:         rds,
		ApplyFriendModel: model.NewApplyFriendModel(db.DB),
		FriendModel:      model.NewFriendModel(db.DB),
	}
}
