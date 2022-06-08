package post

import (
	"BasicOA/response"
	"BasicOA/serve/regular"
	"BasicOA/serve/sector"
	_type "BasicOA/type"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Operate(c *gin.Context) {
	var (
		User     _type.User
		err      error
		Superior _type.Superior
	)
	if User.Email, err = response.GetCookie(c, regular.Email, "email"); err != nil {
		c.JSON(http.StatusOK, _type.Respond{Code: 0, Message: err.Error()})
		return
	}
	if User.Password, err = response.GetCookie(c, regular.Md5, "password"); err != nil {
		c.JSON(http.StatusOK, _type.Respond{Code: 0, Message: err.Error()})
		return
	}
	if _, Superior.User, err = response.PostForm(c, regular.Id, "user"); err != nil {
		c.JSON(http.StatusOK, _type.Respond{Code: 0, Message: err.Error()})
		return
	}
	if _, Superior.License, err = response.PostForm(c, regular.True, "license"); err != nil {
		c.JSON(http.StatusOK, _type.Respond{Code: 0, Message: err.Error()})
		return
	}
	if err = sector.Operate(User, Superior); err != nil {
		c.JSON(http.StatusOK, _type.Respond{Code: 0, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, _type.Respond{Code: 1, Message: "修改成功"})
}
