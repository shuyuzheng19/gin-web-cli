package configs

// ServerConfig 服务器配置
type ServerConfig struct {
	Release      bool   `yaml:"release"`
	Cron         bool   `yaml:"cron"`
	Addr         string `yaml:"addr"`
	ApiPrefix    string `yaml:"apiPrefix"`
	ReadTimeOut  int    `yaml:"readTimeOut"`
	WriteTimeOut int    `yaml:"writeTimeOut"`
}
