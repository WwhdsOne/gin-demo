package initialize

import (
	"fmt"
	"gin-demo/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	m := global.CONFIG.Mysql
	ds := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		m.User,     //用户名
		m.Password, //密码
		m.Host,     //地址
		m.Port,     //端口
		m.Database, //数据库
	)
	// 连接数据库
	Db, err := gorm.Open(mysql.Open(ds), &gorm.Config{})
	if err != nil {
		fmt.Println("连接数据库失败")
		return nil
	}
	return Db
}
