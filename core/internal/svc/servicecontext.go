package svc

import (
	"cherry-disk/core/internal/config"
	"cherry-disk/core/internal/model"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	Db     *gorm.DB
	Rc     *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Db:     model.Init(c),
		Rc:     model.InitRc(c),
	}
}
