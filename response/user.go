package response

type UserResponse struct {
	Id       int    `json:"id"`       //用户ID
	Nickname string `json:"nickname"` //用户名
	Avatar   string `json:"icon"`     //用户头像
	RoleId   uint   `json:"role_id"`  //用户角色
	Username string `json:"username"` //用户账号
}
