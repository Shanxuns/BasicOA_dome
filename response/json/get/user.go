package get

import (
	"BasicOA/response"
	"BasicOA/serve/regular"
	"BasicOA/serve/user"
	_type "BasicOA/type"
	"github.com/gin-gonic/gin"
	"net/http"
)

func User(c *gin.Context) {
	// 创建User结构体
	var (
		User _type.User
		err  error
	)
	// 写入并验证参数
	if User.Email, err = response.GetCookie(c, regular.Email, "email"); err != nil {
		c.JSON(http.StatusOK, _type.Respond{Code: 0, Message: err.Error()})
		return
	}
	if User.Password, err = response.GetCookie(c, regular.Md5, "password"); err != nil {
		c.JSON(http.StatusOK, _type.Respond{Code: 0, Message: err.Error()})
		return
	}
	if err, User = user.User(User); err != nil {
		c.JSON(http.StatusOK, _type.Respond{Code: 0, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, _type.Respond{Code: 1, Message: "获取成功", Data: User})
}
