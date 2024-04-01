package common

type RoleId uint

const (
	//普通用户ID
	USER_ID RoleId = 1
	//管理员ID
	ADMIN_ID RoleId = 2
	//超级管理员ID
	SUPER_ADMIN_ID RoleId = 3
)

const (
	TokenType = "Bearer " //token类型

	TokenHeader = "Authorization" //token请求头
)
