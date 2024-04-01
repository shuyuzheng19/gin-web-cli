package configs

import (
	"os"

	"gopkg.in/yaml.v3"
)

// 全局配置
type GlobalConfig struct {
	IpDB    string       `yaml:"ipDB"`
	Server  ServerConfig `yaml:"server"`
	Author  AuthorInfo   `yaml:"author"`
	Logger  LoggerConfig `yaml:"logger"`
	DB      DbConfig     `yaml:"db"`
	Cors    CorsConfig   `yaml:"cors"`
	Jwt     JwtConfig    `yaml:"jwt"`
	Email   EmailConfig  `yaml:"email"`
	Redis   RedisConfig  `yaml:"redis"`
	MyEmail string       `yaml:"myEmail"`
	Upload  UploadConfig `yaml:"upload"`
}

type AuthorInfo struct {
	Name    string `yaml:"name"`
	Home    string `yaml:"home"`
	Github  string `yaml:"github"`
	Version string `yaml:"version"`
}

type UploadConfig struct {
	MaxImageSize int64  `yaml:"maxImageSize"`
	MaxFileSize  int64  `yaml:"maxFileSize"`
	Uri          string `yaml:"uri"`
	Path         string `yaml:"path"`
}

type CorsConfig struct {
	AllOrigins       bool     `yaml:"allOrigins"`
	Enable           bool     `yaml:"enable"`
	AllowOrigins     []string `yaml:"allowOrigins"`
	AllowMethods     []string `yaml:"allowMethods"`
	AllowHeaders     []string `yaml:"allowHeaders"`
	ExposeHeaders    []string `yaml:"exposeHeaders"`
	AllowCredentials bool     `yaml:"allowCredentials"`
}

type JwtConfig struct {
	Expire  uint   `yaml:"expire"`
	Encrypt string `yaml:"encrypt"`
}

var CONFIG GlobalConfig

func InitConfig(path string) error {
	//读取配置文件
	var bytes, err = os.ReadFile(path)
	if err != nil {
		return err
	}
	//将配置文件的配置解析到结构体中
	err = yaml.Unmarshal(bytes, &CONFIG)
	return err
}
