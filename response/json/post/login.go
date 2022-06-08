package post

import (
	"BasicOA/response"
	"BasicOA/serve/regular"
	"BasicOA/serve/user"
	_type "BasicOA/type"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	// 创建User结构体
	var (
		User _type.User
		err  error
	)
	// 写入并验证参数

	if User.Email, _, err = response.PostForm(c, regular.Email, "email"); err != nil {
		c.JSON(http.StatusOK, _type.Respond{Code: 0, Message: err.Error()})
		return
	}
	if User.Password, _, err = response.PostForm(c, regular.Password, "password"); err != nil {
		c.JSON(http.StatusOK, _type.Respond{Code: 0, Message: err.Error()})
		return
	}
	if err, User = user.Login(User); err != nil {
		c.JSON(http.StatusOK, _type.Respond{Code: 0, Message: err.Error()})
		return
	}
	// 设置cookie
	c.SetCookie("email", User.Email, 3600, "/", "localhost", false, true)
	c.SetCookie("password", User.Password, 3600, "/", "localhost", false, true)
	// 响应结果
	c.JSON(http.StatusOK, _type.Respond{Code: 1, Message: "登录成功", Data: User})
}
