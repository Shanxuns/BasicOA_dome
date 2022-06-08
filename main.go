package main

import (
	"BasicOA/config"
	"BasicOA/serve/email"
	"BasicOA/serve/form"
	"BasicOA/serve/gin"
	"BasicOA/serve/log"
	"BasicOA/serve/mysql"
	"BasicOA/serve/task"
	_type "BasicOA/type"
	"runtime"
	"sync"
)

func main() {
	// 开启全核心
	runtime.GOMAXPROCS(runtime.NumCPU())
	// 日志系统调试模式
	log.SetMode(false)
	// Logo
	var i = " ____            _        ___   __  __ \n| __ )  __ _ ___(_) ___  |_ _| |  \\/  |\n|  _ \\ / _` / __| |/ __|  | |  | |\\/| |\n| |_) | (_| \\__ \\ | (__   | |  | |  | |\n|____/ \\__,_|___/_|\\___| |___| |_|  |_|\n"
	log.Print(i)
	log.Debug("版本:", "1.0.0 图库版")
	wg := sync.WaitGroup{}
	wg.Add(1)
	// 配置文件
	config.Run()
	wg.Add(1)
	// 启动Email服务
	email.RunEmail()
	wg.Add(1)
	// 启动集成Mysql服务
	go mysql.RunMysql()
	// 启动文件服务
	wg.Add(1)
	go _type.Run()
	// 启动表单服务
	wg.Add(1)
	go form.Run()
	// 启动Gin集成服务
	task.Run()
	gin.RunGin()
	wg.Wait()
}
