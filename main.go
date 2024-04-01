package main

import (
	"gin-web/configs"
	"gin-web/helper"
	"gin-web/router"
)

func main() {
	var configPath = "application.yaml"

	var err = configs.InitConfig(configPath)

	helper.PanicError(err)

	var config = configs.CONFIG

	configs.InitIpDBConfig(config.IpDB)

	configs.InitLogger(config.Logger)

	configs.InitDbConfig(config.DB)

	configs.InitRedisConfig(config.Redis)

	router.NewRouter(config.Server).SetupRouters().RunServer()
}
