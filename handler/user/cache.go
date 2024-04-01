package user

import (
	"fmt"
	"gin-web/common"
	"gin-web/configs"
	"gin-web/models"
	"gin-web/request"
	"gin-web/utils"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

type UserCache struct {
	client *redis.Client
}

// SetToken 将token缓存
func (u UserCache) SetToken(id int, token string) error {
	return u.client.Set(fmt.Sprintf(common.UserTokenKey+"%d", id), token, time.Duration(configs.CONFIG.Jwt.Expire)*(time.Hour*24)).Err()
}

// GetToken 从缓存中获取token
func (u UserCache) GetToken(id int) string {
	return u.client.Get(fmt.Sprintf(common.UserTokenKey+"%d", id)).Val()
}

// RemoveToken 删除用户的token
func (u UserCache) RemoveToken(uid int) error {
	return u.client.Del(common.UserTokenKey + strconv.Itoa(uid)).Err()
}

// SetUser 将用户信息缓存
func (u UserCache) SetUser(id int, user *models.User) error {
	var str = utils.ObjectToJson(&user)
	return u.client.Set(common.UserInfoKey+strconv.Itoa(id), str, common.UserInfoKeyExpire).Err()
}

func (u UserCache) DeleteUser(id int) error {
	configs.LOGGER.Info("清除用户信息", zap.Int("uid", id))
	return u.client.Del(common.UserInfoKey + strconv.Itoa(id)).Err()
}

// GetUser 从缓存中获取用户信息
func (u UserCache) GetUser(id int) *models.User {
	var result = u.client.Get(common.UserInfoKey + strconv.Itoa(id)).Val()
	return utils.JsonToObject[models.User](result)
}

// SetEmailCode 将邮箱验证码缓存
func (u UserCache) SetEmailCode(code, email string) error {
	return u.client.Set(fmt.Sprintf(common.EmailCodeKey+"%s", email), code, common.EmailCodeKeyExpire).Err()
}

// GetEmailCode 从缓存中获取邮箱验证码
func (u UserCache) GetEmailCode(email string) string {
	return u.client.Get(fmt.Sprintf(common.EmailCodeKey+"%s", email)).Val()
}

func (u UserCache) SaveRetrieveCode(r request.RetrieveCache) error {
	return u.client.Set(common.Retrieve+r.Email, utils.ObjectToJson[request.RetrieveCache](r), common.RetrieveExpire).Err()
}

func (u UserCache) DeleteRetrieveCode(email string) error {
	return u.client.Del(common.Retrieve + email).Err()
}

func (u UserCache) GetRetrieveCode(email string) *request.RetrieveCache {
	var str = u.client.Get(common.Retrieve + email).Val()
	if str == "" {
		return nil
	} else {
		return utils.JsonToObject[request.RetrieveCache](str)
	}
}

func NewUserCache(client *redis.Client) UserCache {
	return UserCache{client: client}
}
