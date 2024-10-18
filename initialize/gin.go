package initialize

import (
	"github.com/gin-gonic/gin"
)

func InitGin() *gin.Engine {
	// 初始化总路由
	r := gin.Default()
	// 使用gin框架日志
	r.Use(gin.Logger())
	// 使用gin框架的恢复器
	r.Use(gin.Recovery())
	return r
}
