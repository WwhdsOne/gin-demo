package config

// Config 定义配置结构体
type Config struct {
	Server Server `yaml:"server"`
	Mysql  Mysql  `yaml:"mysql"`
	//Zap   Zap   `yaml:"zap"`
}
