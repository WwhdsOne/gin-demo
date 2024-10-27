package service

import (
	"gin-demo/global"
	"gin-demo/model"
)

func Login(u *model.User) *model.User {
	var user model.User
	global.DB.Where("Username = ?", u.Username).Where("Password = ?", u.Password).First(&user)
	return &user
}
