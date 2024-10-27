package initialize

import (
	"github.com/gin-gonic/gin"
)

func InitGin() *gin.Engine {
	// 初始化总路由
	r := gin.Default()
	return r
}
