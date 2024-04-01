package user

import (
	"gin-web/common"
	"gin-web/models"
	"gin-web/request"
	"gin-web/response"
)

type UserRepository interface {
	Create(user models.User) error
	GetById(id int) *models.User
	FindByUsername(username string) *models.User
	FindByUsernameAndEmail(username, email string) *models.User
	UpdatePassword(username, email, newPassword string) int64
	GetUserByPage(req request.UserFilter) (users []response.UserResponse, total int64)
}

type UserService interface {
	//登录
	Login(userRequest request.UserLoginRequest) response.TokenResponse
	//注册
	Register(userRequest request.UserRequest)
	//获取用户信息
	GetUser(id int) *models.User
	//获取用户Token
	GetToken(id int) string
	//找回密码并发送到邮箱
	RetrievePassword(request request.RetrieveRequest)
	//通过邮箱验证码进行修改密码
	UpdatePassword(request request.PasswordRequest)
	//发送注册验证码到邮箱
	SendEmailCode(email string)
	//验证邮箱验证码是否正确
	ValidateEmailCode(email string, code string)
	//联系我
	ContactMe(request request.ContactRequest)
	//获取用户列表
	GetUsers(req request.UserFilter) common.PageInfo
}
