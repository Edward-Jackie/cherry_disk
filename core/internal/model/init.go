package model

import (
	"cherry-disk/core/internal/config"
	"fmt"
	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func Init(c config.Config) *gorm.DB {
	glog := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			Colorful:                  true,
			IgnoreRecordNotFoundError: false,
			LogLevel:                  logger.Info,
		},
	)

	db, err := gorm.Open(
		mysql.Open(c.Mysql.DataSource),
		&gorm.Config{
			Logger: glog,
		})

	if err != nil {
		fmt.Println("gorm 连接数据库失败")
	}

	return db
}

func InitRc(c config.Config) *redis.Client {
	client := redis.NewClient(
		&redis.Options{
			Addr:     c.Redis.Addr,
			Password: c.Redis.Password,
			DB:       c.Redis.DB,
		},
	)

	_, err := client.Ping().Result()
	if err != nil {
		fmt.Println("Redis 连接失败！！ err = ", err)
		return nil
	}
	return client
}
