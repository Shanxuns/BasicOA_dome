package mysql

import (
	"BasicOA/config"
	"BasicOA/serve/log"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func RunMysql() {
	var appName = "Mysql"
	if Db != nil {
		log.App(appName, "服务已经启动")
	} else {
		log.App(appName, "启动服务")
		USERNAME, err := config.Conf.GetValue("mysql", "USERNAME")
		if err != nil {
			log.Err(appName, "用户名配置读取失败", err)
			return
		}
		PASSWORD, err := config.Conf.GetValue("mysql", "PASSWORD")
		if err != nil {
			log.Err(appName, "密码配置读取失败", err)
			return
		}
		NETWORK, err := config.Conf.GetValue("mysql", "NETWORK")
		if err != nil {
			log.Err(appName, "协议配置读取失败", err)
			return
		}
		SERVER, err := config.Conf.GetValue("mysql", "SERVER")
		if err != nil {
			log.Err(appName, "IP配置读取失败", err)
			return
		}
		PORT, err := config.Conf.GetValue("mysql", "PORT")
		if err != nil {
			log.Err(appName, "端口配置读取失败", err)
			return
		}
		DATABASE, err := config.Conf.GetValue("mysql", "DATABASE")
		if err != nil {
			log.Err(appName, "数据库名配置读取失败", err)
			return
		}
		// 建立数据库连接
		Db, errDb = sql.Open("mysql", fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=utf8", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE))
		if errDb != nil {
			log.Err(appName, "数据库连接失败", errDb)
			return
		}
		// 数据库最大连接周期，永不超时超过时间
		Db.SetConnMaxLifetime(0)
		// 数据库设置最大连接数
		Db.SetMaxOpenConns(0)
		// 数据库设置闲置连接数
		Db.SetMaxIdleConns(0)
		// 数据库Ping测试
		err = Db.Ping()
		if err != nil {
			log.Err(appName, "数据库Ping测试失败", err)
			return
		}
	}
}
