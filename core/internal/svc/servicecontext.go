package svc

import (
	"cherry-disk/core/internal/config"
	"cherry-disk/core/internal/middleware"
	"cherry-disk/core/internal/model"
	"github.com/go-redis/redis"
	"github.com/zeromicro/go-zero/rest"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	Db     *gorm.DB
	Rc     *redis.Client
	Auth   rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Db:     model.Init(c),
		Rc:     model.InitRc(c),
		Auth:   middleware.NewAuthMiddleWare().Handle,
	}
}
