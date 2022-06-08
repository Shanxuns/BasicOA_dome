package post

import (
	"BasicOA/response"
	"BasicOA/serve/regular"
	"BasicOA/serve/user"
	_type "BasicOA/type"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	// 创建User结构体
	var (
		User   _type.User
		err    error
		Avatar string
	)
	// 写入并验证参数
	if User.Email, _, err = response.PostForm(c, regular.Email, "email"); err != nil {
		c.JSON(http.StatusOK, _type.Respond{Code: 0, Message: err.Error()})
		return
	}
	if User.Fname, _, err = response.PostForm(c, regular.Fname, "fname"); err != nil {
		c.JSON(http.StatusOK, _type.Respond{Code: 0, Message: err.Error()})
		return
	}
	if User.Lname, _, err = response.PostForm(c, regular.Lname, "lname"); err != nil {
		c.JSON(http.StatusOK, _type.Respond{Code: 0, Message: err.Error()})
		return
	}
	if User.Phone, _, err = response.PostForm(c, regular.Phone, "phone"); err != nil {
		c.JSON(http.StatusOK, _type.Respond{Code: 0, Message: err.Error()})
		return
	}
	if User.Password, _, err = response.PostForm(c, regular.Password, "password"); err != nil {
		c.JSON(http.StatusOK, _type.Respond{Code: 0, Message: err.Error()})
		return
	}
	User.Avatar = "default"
	if Avatar, _, err = response.PostForm(c, regular.Avatar, "avatar"); err == nil {
		User.Avatar = Avatar
	}
	if err = user.Register(User); err != nil {
		c.JSON(http.StatusOK, _type.Respond{Code: 0, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, _type.Respond{Code: 1, Message: "注册成功"})
}
