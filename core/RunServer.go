package core

import (
	"gin-demo/global"
	"gin-demo/initialize"
	"net/http"
	"time"
)

func RunServer() {
	// 初始化数据库
	global.DB = initialize.InitDB()
	// 初始化GIN
	gin := initialize.InitGin()

	// 服务器设置
	serverConfig := global.CONFIG.Server
	s := &http.Server{
		Addr:           serverConfig.Port,
		Handler:        gin,
		ReadTimeout:    time.Duration(serverConfig.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(serverConfig.WriteTimeout) * time.Second,
		MaxHeaderBytes: int(serverConfig.MaxHeaderBytes),
	}
	s.ListenAndServe()
}
