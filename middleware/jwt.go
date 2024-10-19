package middleware

import (
	"gin-demo/pkg/response"
	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("w-token")
		if token == "" {
			response.NoAuth("未登录或非法访问", c)
			c.Abort() // 终止请求
			return
		}
	}
}
