package common

// R 全局返回
type R struct {
	Code    int         `json:"code"`           //状态码
	Message string      `json:"message"`        //返回的消息
	Data    interface{} `json:"data,omitempty"` //返回的数据
	Error   error       `json:"-"`
}

type PageInfo struct {
	Page  int         `json:"page"`
	Size  int         `json:"size"`
	Total int64       `json:"total"`
	Data  interface{} `json:"data"`
}

// 错误代码
type ErrorCode uint

// 返回成功
func OK() R {
	return R{Code: 200, Message: "成功"}
}

// 成功并返回数据
func SUCCESS(data interface{}) R {
	var r = OK()
	r.Data = data
	return r
}

const (
	ERROR        ErrorCode = 500
	BadRequest   ErrorCode = 100
	Unauthorized ErrorCode = 401
	Forbidden    ErrorCode = 403
	NotFound     ErrorCode = 404
	FAIL                   = iota + 10001
	LoginFail
	CreateTokenFail
	RegisteredCode
	SendEmailFailCode
	NoLogin
	ParseTokenFail
	TokenExpireFail
	UpdateFail
	EmailCodeValidate
	NoFile
)

// 错误信息集合
var errors = map[ErrorCode]string{
	ERROR:             "服务器错误",
	FAIL:              "处理失败",
	BadRequest:        "参数验证失败",
	LoginFail:         "账号或密码错误",
	CreateTokenFail:   "生成Token失败",
	RegisteredCode:    "注册账号失败",
	SendEmailFailCode: "发送邮件失败",
	NoLogin:           "还未登录",
	ParseTokenFail:    "错误的token",
	TokenExpireFail:   "Token已过期",
	Unauthorized:      "未授权",
	Forbidden:         "你没有权限访问",
	NotFound:          "找不到相关信息",
	UpdateFail:        "修改失败",
	EmailCodeValidate: "错误的邮箱验证码",
	NoFile:            "没有文件",
}

type E struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
	Error   error     `json:"-"`
}

// 返回失败
func AutoFail(code ErrorCode) E {
	if message, ok := errors[code]; ok {
		return E{Code: code, Message: message}
	} else {
		return AutoFail(FAIL)
	}
}

func BadRequestFail(message string) E {
	return E{Code: BadRequest, Message: message}
}
