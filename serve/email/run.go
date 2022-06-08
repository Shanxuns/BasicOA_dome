package email

import (
	"BasicOA/config"
	"BasicOA/serve/log"
)

func RunEmail() {
	var (
		err     error
		appName = "Email"
	)
	log.App(appName, "启动服务")
	SMTPHost, err = config.Conf.GetValue("Email", "SMTPHost")
	if err != nil {
		log.Err(appName, "域名配置读取失败", err)
		return
	}
	SMTPPort, err = config.Conf.GetValue("Email", "SMTPPort")
	if err != nil {
		log.Err(appName, "端口配置读取失败", err)
		return
	}
	SMTPUsername, err = config.Conf.GetValue("Email", "SMTPUsername")
	if err != nil {
		log.Err(appName, "用户名配置读取失败", err)
		return
	}
	SMTPPassword, err = config.Conf.GetValue("Email", "SMTPPassword")
	if err != nil {
		log.Err(appName, "密码配置读取失败", err)
		return
	}
}
