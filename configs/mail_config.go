package configs

import (
	"fmt"
	"gin-web/common"
	"gin-web/helper"
	"net/smtp"

	"github.com/jordan-wright/email"
)

type EmailConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Addr     string `yaml:"addr"`
}

// SendEmail 发送邮件
// to:对方邮箱 subject:邮箱主题 isHTML:是否是html格式 text:文本信息
func (ef EmailConfig) SendEmail(to string, subject string, isHTML bool, text string) {

	var e = email.NewEmail()

	e.From = fmt.Sprintf("%s <%s>", "", ef.Username)

	e.To = []string{to}

	e.Subject = subject

	if isHTML {
		e.HTML = []byte(text)
	} else {
		e.Text = []byte(text)
	}

	if err := e.Send(ef.Addr, smtp.PlainAuth("", ef.Username, ef.Password, ef.Host)); err != nil {
		helper.ErrorToResponseAndError(common.SendEmailFailCode, err)
	}
}
