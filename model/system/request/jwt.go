package request

import (
	"github.com/gofrs/uuid/v5"
	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	BaseClaims                 // 基础claims
	BufferTime           int64 // 缓冲时间
	jwt.RegisteredClaims       // 注册claims
}

type BaseClaims struct {
	UUID        uuid.UUID // 唯一标识uuid
	ID          int64     // 主键id
	Username    uint64    // 用户名
	NickName    string    // 用户昵称
	AuthorityId uint      // 角色ID
}
