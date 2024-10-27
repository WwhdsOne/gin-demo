package model

import "github.com/gofrs/uuid/v5"

type User struct {
	BaseModel
	Username    uint64    // 用户账号
	UUID        uuid.UUID // uuid
	Password    string    // 用户密码
	Nickname    string    // 用户昵称
	AuthorityId uint      // 角色
}
