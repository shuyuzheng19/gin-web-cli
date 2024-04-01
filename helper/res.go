package helper

import (
	"gin-web/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 处理成功,返回数据到前端
func ResultSuccessToResponse(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, common.SUCCESS(data))
}

// 处理失败,返回错误信息到前端
func ResultFailToResponse(ctx *gin.Context, code common.ErrorCode) {
	ctx.AbortWithStatusJSON(http.StatusOK, common.AutoFail(code))
}

// 处理失败,返回错误信息到前端
func ResultBadRequestFail(ctx *gin.Context, info common.E) {
	ctx.AbortWithStatusJSON(http.StatusOK, info)
}
