package utils

import (
	"github.com/gin-gonic/gin"
)

// GetToken 从header中获取x-token
func GetToken(c *gin.Context) string {
	token := c.GetHeader("x-token")
	return token
}
