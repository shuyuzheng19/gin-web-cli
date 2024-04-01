package middleware

import (
	"gin-web/common"
	"gin-web/helper"
	"gin-web/models"
	"gin-web/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

var GetJwtUser = func(id int) *models.User {
	return &models.User{}
}

var GetToken = func(id int) string {
	return ""
}

func ParseToken(header string, context *gin.Context) *models.User {
	if header == "" || !strings.HasPrefix(header, common.TokenType) {
		helper.ResultFailToResponse(context, common.NoLogin)
		return nil
	}

	var token = strings.Replace(header, common.TokenType, "", 1)

	var uid = utils.ParseTokenToUserId(token)

	if uid == -1 {
		helper.ResultFailToResponse(context, common.ParseTokenFail)
		return nil
	}

	var redisToken = GetToken(uid)

	if redisToken != token {
		helper.ResultFailToResponse(context, common.TokenExpireFail)
		return nil
	}

	var user = GetJwtUser(uid)

	if user.ID == 0 {
		helper.ResultFailToResponse(context, common.Unauthorized)
		return nil
	}

	return user
}

// JwtMiddle 验证身份中间件
func JwtMiddle(roleId common.RoleId) gin.HandlerFunc {
	return func(context *gin.Context) {
		var header = context.GetHeader(common.TokenHeader)

		var user = ParseToken(header, context)

		if context.IsAborted() {
			return
		}

		var role = user.Role.ID

		var isAuth = false

		if roleId == common.USER_ID || role == uint(common.SUPER_ADMIN_ID) {
			isAuth = true
		} else if roleId == common.ADMIN_ID && role == uint(common.ADMIN_ID) {
			isAuth = true
		}

		if isAuth {
			context.Set("user", user)
			context.Next()
		} else {
			helper.ResultFailToResponse(context, common.Forbidden)
			return
		}

	}
}
