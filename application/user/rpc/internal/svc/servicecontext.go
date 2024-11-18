package svc

import (
	"im/application/user/rpc/internal/config"
	"im/application/user/rpc/internal/model"
	"im/pkg/orm"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config    config.Config
	DB        *orm.DB
	BizRedis  *redis.Redis
	UserModel *model.UserModel
	// EmailRpc  email.Email
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
		Config:    c,
		DB:        db,
		BizRedis:  rds,
		UserModel: model.NewUserModel(db.DB),
		// EmailRpc:  email.NewEmail(zrpc.MustNewClient(c.EmailRpc)),
	}
}
