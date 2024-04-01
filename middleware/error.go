package middleware

import (
	"gin-web/common"
	"gin-web/configs"
	"gin-web/helper"
	"gin-web/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ErrorMiddle 全局错误中间件
func ErrorMiddle(ctx *gin.Context) {
	defer func() {
		var err interface{} = recover()

		if err != nil {
			switch t := err.(type) {
			case common.E:
				var ip = utils.GetIPAddress(ctx.Request)
				var city = utils.GetIpCity(ip)
				configs.LOGGER.Warn("自定义错误",
					zap.String("path", ctx.FullPath()),
					zap.Any("code", t.Code),
					zap.String("message", t.Message),
					zap.String("method", ctx.Request.Method),
					zap.String("ip", ip),
					zap.String("city", city),
					zap.Any("error", t.Error))
				ctx.AbortWithStatusJSON(200, t)
			default:
				configs.LOGGER.Warn("服务器抛出错误",
					zap.String("path", ctx.FullPath()),
					zap.String("method", ctx.Request.Method), zap.String("error", err.(error).Error()))
				helper.ResultFailToResponse(ctx, common.ERROR)
			}
		}
	}()

	ctx.Next()
}
