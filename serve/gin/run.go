package gin

import (
	"BasicOA/config"
	"BasicOA/serve/log"
	"github.com/gin-gonic/gin"
)

func RunGin() {
	var appName = "Gin"
	NETWORK, NETWORKErr := config.Conf.GetValue("BasicIM", "NETWORK")
	CertPath, CertErr := config.Conf.GetValue("BasicIM", "CertPath")
	KeyPath, KeyErr := config.Conf.GetValue("BasicIM", "KeyPath")
	Port, PortErr := config.Conf.GetValue("BasicIM", "PORT")
	DEBUG, DEBUGErr := config.Conf.GetValue("BasicIM", "DEBUG")
	if DEBUGErr == nil {
		// 是否开启调试模式
		if DEBUG == "false" {
			// 关闭调试模式
			gin.SetMode(gin.ReleaseMode)
		}
	}
	// 创建Gin服务
	Gin = gin.Default()
	Gin.Use(IsMysql)
	// 加载路由控制
	runRoute()
	// 运行Gin服务
	if PortErr == nil {
		if NETWORKErr != nil {
			log.App(appName, "启动服务 http://127.0.0.1:"+Port)
			err := Gin.Run(":" + Port)
			if err != nil {
				// 错误处理
				log.Err(appName, "启动失败", err)
				return
			}
		} else if NETWORK == "HTTPS" {
			Gin.Use(TlsHandler(Port))
			log.App(appName, "启动服务 https://127.0.0.1:"+Port)
			if CertErr != nil || KeyErr != nil {
				log.Err(appName, "启动失败", "获取证书失败")
			}
			err := Gin.RunTLS(":"+Port, CertPath, KeyPath)
			if err != nil {
				// 错误处理
				log.Err(appName, "启动失败", err)
				return
			}
		} else {
			log.App(appName, "启动服务 http://127.0.0.1:"+Port)
			err := Gin.Run(":" + Port)
			if err != nil {
				// 错误处理
				log.Err(appName, "启动失败", err)
				return
			}
		}
		return
	}
	if NETWORKErr != nil {
		log.App(appName, "启动服务 http://127.0.0.1:8080")
		err := Gin.Run(":8080")
		if err != nil {
			// 错误处理
			log.Err(appName, "启动失败", err)
			return
		}
	} else if NETWORK == "HTTPS" {
		Gin.Use(TlsHandler("8080"))
		log.App(appName, "启动服务 https://127.0.0.1:8080")
		err := Gin.RunTLS(":8080", "", "")
		if err != nil {
			// 错误处理
			log.Err(appName, "启动失败", err)
			return
		}
	} else {
		log.App(appName, "启动服务 http://127.0.0.1:8080")
		err := Gin.Run(":8080")
		if err != nil {
			// 错误处理
			log.Err(appName, "启动失败", err)
			return
		}
	}
	return
}
