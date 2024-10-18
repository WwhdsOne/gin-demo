package config

type App struct {
	Name          string `yaml:"name"`
	Version       string `yaml:"version"`
	Port          int    `yaml:"port"`
	Debug         bool   `yaml:"debug"`
	CaseSensitive bool   `yaml:"caseSensitive"`
	StrictRouting bool   `yaml:"strictRouting"`
	ServerHeader  string `yaml:"serverHeader"`
	AppName       string `yaml:"appName"`
}
