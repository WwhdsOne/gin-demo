package config

// Config 定义配置结构体
type Config struct {
	App   App   `yaml:"app"`
	Mysql Mysql `yaml:"mysql"`
	//Zap   Zap   `yaml:"zap"`
}
