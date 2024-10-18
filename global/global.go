package global

import (
	"gin-demo/initialize/config"
	"gorm.io/gorm"
)

var (
	// CONFIG 服务配置
	CONFIG *config.Config
	// DB 数据库连接
	DB *gorm.DB
	// LOG 日志
	//LOG *zap.Logger
)
