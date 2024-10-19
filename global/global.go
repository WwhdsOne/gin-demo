package global

import (
	"gin-demo/initialize/config"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	// CONFIG 服务配置
	CONFIG *config.Config
	// DB 数据库连接
	DB *gorm.DB
	// Redis redis连接
	Redis *redis.Client
	// LOG 日志
	//LOG *zap.Logger
)
