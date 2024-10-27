package core

import (
	"gin-demo/global"
	"gin-demo/initialize"
)

func RunServer() {
	// 初始化数据库
	global.DB = initialize.InitDB()
	// 初始化Redis
	global.Redis = initialize.InitRedis()
	// 初始化GIN
	gin := initialize.InitGin()

	// 初始化路由
	initialize.InitRouter(gin)

	// 服务器设置
	//serverConfig := global.CONFIG.Server
	//s := &http.Server{
	//	Addr:           serverConfig.Port,
	//	Handler:        gin,
	//	ReadTimeout:    time.Duration(serverConfig.ReadTimeout) * time.Second,
	//	WriteTimeout:   time.Duration(serverConfig.WriteTimeout) * time.Second,
	//	MaxHeaderBytes: int(serverConfig.MaxHeaderBytes),
	//}
	//s.ListenAndServe()
	// 启动
	gin.Run(global.CONFIG.Server.Port)
}
