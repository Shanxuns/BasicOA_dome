package email

import (
	"BasicOA/serve/log"
	"fmt"
	"net/smtp"
	"time"
)

var (
	SMTPHost     = "smtp.mxhichina.com"
	SMTPPort     = ":25"
	SMTPUsername = "shanxun@kuocaitm.net"
	SMTPPassword = "Tanjicai@qq.com"
	MaxClient    = 5
)

var pool *Pool

func SendEmail(receiver string, subject string, text string) {
	var err error
	var appName = "Email"
	if pool == nil {
		pool, err = NewPool(SMTPHost+SMTPPort, MaxClient, smtp.PlainAuth("", SMTPUsername, SMTPPassword, SMTPHost))
		if err != nil {
			log.Err(appName, "配置错误", err)
		}
	}
	e := &Email{
		From:    fmt.Sprintf("BasicIM 官方账号 <%s>", SMTPUsername),
		To:      []string{receiver},
		Subject: subject,
		HTML:    []byte(text),
	}
	err = pool.Send(e, 5*time.Second)
	if err != nil {
		log.Err(appName, "发送失败"+SMTPHost+SMTPPort+" "+SMTPUsername, err)
	}
}
