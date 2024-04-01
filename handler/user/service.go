package user

import (
	"fmt"
	"gin-web/common"
	"gin-web/configs"
	"gin-web/helper"
	"gin-web/models"
	"gin-web/request"
	"gin-web/response"
	"gin-web/utils"

	"go.uber.org/zap"
)

type UserServiceImpl struct {
	repository UserRepository
	cache      UserCache
}

func (u UserServiceImpl) Login(userRequest request.UserLoginRequest) response.TokenResponse {
	var user = u.repository.FindByUsername(userRequest.Username)
	if user == nil {
		helper.ErrorToResponse(common.LoginFail)
	}

	if !utils.ValidatorPassword(userRequest.Password, user.Password) {
		helper.ErrorToResponse(common.LoginFail)
	}

	var token = utils.CreateAccessToken(user.ID, user.Username)

	go func() {
		u.cache.SetToken(user.ID, token.Token)
		u.cache.SetUser(user.ID, user)
	}()

	configs.LOGGER.Info("用户登录成功", zap.Int("id", user.ID),
		zap.String("username", user.Username),
		zap.String("nickname", user.NickName))

	return token
}

func (u UserServiceImpl) Register(userRequest request.UserRequest) {

	u.ValidateEmailCode(userRequest.Email, userRequest.Code)

	var user = userRequest.ToUserDo()

	if err := u.repository.Create(user); err != nil {
		helper.ErrorToResponseAndError(common.RegisteredCode, err)
	}

	configs.LOGGER.Info("用户注册成功", zap.Int("id", user.ID),
		zap.String("username", user.Username),
		zap.String("nickname", user.NickName))
}

func (u UserServiceImpl) GetUser(id int) *models.User {
	var user = u.cache.GetUser(id)
	if user == nil {
		user = u.repository.GetById(id)
		if user != nil {
			u.cache.SetUser(user.ID, user)
		} else {
			u.cache.SetUser(id, nil)
		}
	}
	return user
}

func (u UserServiceImpl) RetrievePassword(r request.RetrieveRequest) {
	var user = u.repository.FindByUsernameAndEmail(r.Username, r.Email)

	if user == nil {
		helper.ErrorToResponse(common.NotFound)
	}

	var code = utils.RandomNumberCode()

	configs.CONFIG.Email.SendEmail(r.Email, "找回密码验证码", false, fmt.Sprintf("你的验证码为:%s 请在5分钟内进行操作", code))

	u.cache.SaveRetrieveCode(request.RetrieveCache{Id: user.ID, Username: user.Username, Email: user.Email, Code: code})
}

func (u UserServiceImpl) UpdatePassword(request request.PasswordRequest) {

	var data = u.cache.GetRetrieveCode(request.Email)

	if data == nil {
		helper.ErrorToResponse(common.NotFound)
	}

	var bcryptPassword = utils.BcryptPassword(request.Password)

	if count := u.repository.UpdatePassword(data.Username, request.Email, bcryptPassword); count == 0 {
		helper.ErrorToResponse(common.UpdateFail)
	}

	go func() {
		u.cache.DeleteUser(data.Id)
		u.cache.DeleteRetrieveCode(data.Email)
	}()

	configs.LOGGER.Info("用户修改密码", zap.Int("id", data.Id), zap.String("email", data.Email), zap.String("username", data.Username))
}

func (u UserServiceImpl) GetToken(id int) string {
	return u.cache.GetToken(id)
}

func (u UserServiceImpl) ValidateEmailCode(email string, code string) {

	var cacheCode = u.cache.GetEmailCode(email)

	if cacheCode == "" || cacheCode != code {
		helper.ErrorToResponse(common.EmailCodeValidate)
	}
}

func (u UserServiceImpl) SendEmailCode(email string) {
	var code = utils.RandomNumberCode()

	configs.CONFIG.Email.SendEmail(email, "注册验证码", false, code)

	configs.LOGGER.Info("发送邮箱验证码", zap.String("code", code))

	u.cache.SetEmailCode(code, email)
}

func (u UserServiceImpl) ContactMe(req request.ContactRequest) {
	var text = fmt.Sprintf("<h3>%s</h3><p>对方名字: %s</p><p>对方邮箱: %s</p>留言内容:<p>%s</p>",
		req.Subject, req.Name, req.Email, req.Content)

	configs.CONFIG.Email.SendEmail(configs.CONFIG.MyEmail, req.Subject, true, text)

	configs.LOGGER.Info("联系我，用户发送信息到我的邮件，请注意查收", zap.Any("info", req))
}

func (u UserServiceImpl) GetUsers(req request.UserFilter) common.PageInfo {
	var blogs, total = u.repository.GetUserByPage(req)

	return common.PageInfo{
		Page:  req.Page,
		Size:  common.UserPageSize,
		Total: total,
		Data:  blogs,
	}
}

func NewUserService() UserService {
	return UserServiceImpl{repository: NewUserRepository(configs.DB), cache: NewUserCache(configs.REDIS)}
}
