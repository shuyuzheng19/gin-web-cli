package common

import "time"

// 用户缓存键集合
const (
	UserTokenKey       = "USER-TOKEN:"    //缓存用户Token的Key
	EmailCodeKey       = "EMAIL-CODE:"    //缓存注册邮箱验证码的key
	EmailCodeKeyExpire = time.Minute * 1  //邮箱验证码过期实际
	UserInfoKey        = "USER-INFO:"     //缓存用户信息的key
	UserInfoKeyExpire  = time.Minute * 30 //用户信息过期时间
	Retrieve           = "RETRIEVE_CODE:" //忘记密码，存取的验证码
	RetrieveExpire     = time.Minute * 5  //验证码的过期时间
)
