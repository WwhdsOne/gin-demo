package core

import (
	"gin-demo/global"
	"gin-demo/initialize"
)

func RunServer() {
	// 初始化数据库
	global.DB = initialize.InitDB()
	// 初始化路由
	gin := initialize.InitGin()

	gin.Run(":8080")
}
