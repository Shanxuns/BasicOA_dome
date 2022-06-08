package gin

import (
	"BasicOA/serve/log"
	"BasicOA/serve/mysql"
	_type "BasicOA/type"
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
	"net/http"
)

var (
	Gin *gin.Engine
)

// CloseMysql 结束Mysql服务
func CloseMysql() bool {
	if mysql.Db != nil {
		err := mysql.Db.Close()
		if err != nil {
			return false
		}
	}
	mysql.Db = nil
	mysql.RunMysql()
	return true
}

// IsMysql 验证并重启Mysql服务
func IsMysql(c *gin.Context) {
	if mysql.Db == nil {
		log.Err("Mysql", "重启Mysql服务", "Ping验证不通过")
		CloseMysql()
		c.JSON(http.StatusOK, _type.Respond{Code: 0, Message: "Mysql服务未启动"})
		return
	}
	if mysql.Db.Ping() != nil {
		log.Err("Mysql", "重启Mysql服务", "Ping验证不通过")
		CloseMysql()
		c.JSON(http.StatusOK, _type.Respond{Code: 0, Message: "Mysql服务未启动"})
		return
	}
	c.Header("Access-Control-Allow-Origin", "*")
}

// TlsHandler HTTPS 服务
func TlsHandler(Port string) gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     "localhost:" + Port,
		})
		err := secureMiddleware.Process(c.Writer, c.Request)
		// If there was an error, do not continue.
		if err != nil {
			return
		}
		c.Next()
	}
}
