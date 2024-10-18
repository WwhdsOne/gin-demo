package config

type Server struct {
	ReadTimeout    int    `yaml:"readTimeout"`
	WriteTimeout   int    `yaml:"writeTimeout"`
	Port           string `yaml:"port"`
	MaxHeaderBytes int    `yaml:"maxHeaderBytes"`
}
