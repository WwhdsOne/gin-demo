package controller

import (
	"gin-demo/model"
	"gin-demo/model/system/request"
	"gin-demo/pkg/response"
	"gin-demo/pkg/utils"
	"gin-demo/service"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(r *gin.RouterGroup) {
	// 路由组
	{
		r.POST("/login", Login)
		r.DELETE("/logout", Logout)
	}
}

func Login(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if user.Username == 0 || user.Password == "" {
		response.FailWithMessage("用户名或密码不能为空", c)
		return
	}
	u := service.Login(&user)
	if u == nil {
		response.FailWithMessage("用户名或密码错误", c)
		return
	}
	j := utils.NewJWT() // 唯一签名
	claim := j.CreateClaims(
		request.BaseClaims{
			ID:          u.ID,
			Username:    u.Username,
			UUID:        u.UUID,
			AuthorityId: u.AuthorityId,
		})
	token, err := j.CreateToken(claim)
	if err != nil {
		response.FailWithMessage("生成Token失败", c)
		return
	}
	response.OkWithData(token, c)
}

func Logout(c *gin.Context) {
}
