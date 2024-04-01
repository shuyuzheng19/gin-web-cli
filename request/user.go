package request

import (
	"gin-web/common"
	"gin-web/models"
	"gin-web/utils"
)

type UserLoginRequest struct {
	Username string `json:"username" binding:"required" msg:"账号不能为空"`
	Password string `json:"password" binding:"required" msg:"密码不能为空"`
}

type UserRequest struct {
	Username string `json:"username" binding:"required,min=8,max=16" msg:"账号要大于8个字符,并且不能超过16个字符"`
	Password string `json:"password" binding:"required,min=8,max=16" msg:"密码要大于8个字符,并且不能超过16个字符"`
	Email    string `json:"email" binding:"required,email" msg:"这不是一个邮箱,请检查邮箱格式是否正确"`
	NickName string `json:"nickName" binding:"required,max=50,min=1" msg:"用户名称最低要有1个字符,并且不能超过50个字符"`
	Code     string `json:"code" binding:"required,min=6,max=6" msg:"请输入6位的验证码"`
}

type SendEmailCode struct {
	Email string `form:"email" binding:"required,email" msg:"这不是一个邮箱,请检查邮箱格式是否正确"`
}

type RetrieveRequest struct {
	Username string `json:"username" binding:"required" msg:"账号不能为空"`
	Email    string `json:"email" binding:"required,email" msg:"这不是一个正确的邮箱格式"`
}

type RetrieveCache struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Code     string `json:"code"`
}

type PasswordRequest struct {
	Email    string `json:"email" binding:"required,email" msg:"这不是一个正确的邮箱格式"`
	Password string `json:"password" binding:"required,min=8,max=16" msg:"新密码要大于8个字符,并且不能超过16个字符"`
	Code     string `json:"code" binding:"required,min=6,max=6" msg:"请输入6位的验证码"`
}

type ContactRequest struct {
	Name    string `json:"name"  binding:"required" msg:"名字不能为空"`
	Email   string `json:"email" binding:"required,email" msg:"这不是一个正确的邮箱格式"`
	Subject string `json:"subject" binding:"required,max=50" msg:"主题不能为空并且不能大于50个字符"`
	Content string `json:"content" binding:"required" msg:"请输入主题内容"`
}

type UserSort string

const (
	CREATE UserSort = "CREATE" //通过创建日期排序
	UPDATE UserSort = "UPDATE" //通过修改日期排序
	ROLE   UserSort = "ROLE"   //通过角色排序
)

func (sort UserSort) GetOrderString(prefix string) string {
	switch sort {
	case CREATE:
		return prefix + "created_at desc"
	case UPDATE:
		return prefix + "updated_at  desc"
	case ROLE:
		return prefix + "role_id desc"
	default:
		return prefix + "created_at desc"
	}
}

type UserFilter struct {
	Page   int      `form:"page"`
	RoleID int      `form:"rid"`
	Sort   UserSort `form:"sort"`
}

func (r UserRequest) ToUserDo() models.User {
	return models.User{
		Username: r.Username,
		Password: utils.BcryptPassword(r.Password),
		Email:    r.Email,
		NickName: r.NickName,
		RoleID:   uint(common.USER_ID),
	}
}
