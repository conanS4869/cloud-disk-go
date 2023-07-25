package svc

import (
	"cloud-disk-go/core/internal/config"
	"cloud-disk-go/core/internal/middleware"
	"cloud-disk-go/core/models"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/rest"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config
	Engine *xorm.Engine
	RDB    *redis.Client
	Auth   rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Engine: models.Init(c.MySql.DataSource),
		RDB:    models.InitRedis(c),
		Auth:   middleware.NewAuthMiddleware().Handle,
	}
}
