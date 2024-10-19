package config

type Redis struct {
	Addr     string `yaml:"addr"`     // 地址
	Port     string `yaml:"port"`     // 端口
	Password string `yaml:"password"` // 密码（如果没有密码则为空）
	DB       int    `yaml:"db"`       // 数据库编号
	Protocol int    `yaml:"protocol"` // 指定使用 RESP 3 协议
}
