package middleware

import (
	"gin-web/configs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// 全局跨域
func Cors(corsConfig configs.CorsConfig) gin.HandlerFunc {
	var cConfig = cors.Config{
		AllowMethods:     corsConfig.AllowMethods,
		AllowHeaders:     corsConfig.AllowHeaders,
		AllowCredentials: corsConfig.AllowCredentials,
		ExposeHeaders:    corsConfig.ExposeHeaders,
	}

	if corsConfig.AllOrigins {
		cConfig.AllowAllOrigins = true
	} else {
		cConfig.AllowOrigins = corsConfig.AllowOrigins
	}

	return cors.New(cConfig)
}
