package config

import (
	"BasicOA/serve/log"
	"github.com/Unknwon/goconfig"
)

func Run() {
	var err interface{}
	Conf, err = goconfig.LoadConfigFile("conf.ini")
	if err != nil {
		log.Err("配置系统", "读取失败", err)
		return
	}
}
