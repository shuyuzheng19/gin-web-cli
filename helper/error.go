package helper

import (
	"gin-web/common"
	"log"
)

// 抛出错误
func PanicError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// 抛出错误,并且追加自定义的错误信息
func PanicErrorAndMessage(err error, message string) {
	if err != nil {
		log.Fatal(err, "\t", message)
	}
}

// ErrorToResponse 全局error处理 返回给前端
func ErrorToResponse(errCode common.ErrorCode) {
	panic(common.AutoFail(errCode))
}

// ErrorToResponse 全局error处理 返回给前端
func ErrorToResponseAndError(errCode common.ErrorCode, err error) {
	var r = common.AutoFail(errCode)
	r.Error = err
	panic(r)
}
