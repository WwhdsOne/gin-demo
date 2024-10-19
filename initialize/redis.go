package initialize

import (
	"fmt"
	"gin-demo/global"
	"github.com/redis/go-redis/v9"
)

func InitRedis() *redis.Client {
	r := global.CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", r.Addr, r.Port),
		Password: r.Password, // no password set
		DB:       0,          // use default DB
		Protocol: r.Protocol,
	})
	// 自动迁移
	return client
}
