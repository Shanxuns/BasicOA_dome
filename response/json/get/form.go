package get

import (
	"BasicOA/response"
	"BasicOA/serve/form"
	"BasicOA/serve/regular"
	_type "BasicOA/type"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TypeList(c *gin.Context) {
	var (
		Form int
		err  error
	)
	if _, Form, err = response.GetParam(c, regular.Id, "id"); err != nil {
		c.JSON(http.StatusOK, _type.Respond{Code: 0, Message: err.Error()})
		return
	}
	err, list := form.Db(Form)
	if err != nil {
		c.JSON(http.StatusOK, _type.Respond{Code: 0, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, _type.Respond{Code: 1, Message: "加载表单结构列表成功", Data: list})
}

func List(c *gin.Context) {
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
	err, list := form.List(User)
	if err != nil {
		c.JSON(http.StatusOK, _type.Respond{Code: 0, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, _type.Respond{Code: 1, Message: "加载表单列表成功", Data: list})
}

func Title(c *gin.Context) {
	var (
		Form int
		err  error
	)
	if _, Form, err = response.GetParam(c, regular.Id, "id"); err != nil {
		c.JSON(http.StatusOK, _type.Respond{Code: 0, Message: err.Error()})
		return
	}
	err, FormAdmin := form.Title(Form)
	if err != nil {
		c.JSON(http.StatusOK, _type.Respond{Code: 0, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, _type.Respond{Code: 1, Message: "获取表单信息成功", Data: FormAdmin})
}

func Sqlite(c *gin.Context) {
	var (
		Form int
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
	if _, Form, err = response.GetParam(c, regular.Id, "id"); err != nil {
		c.JSON(http.StatusOK, _type.Respond{Code: 0, Message: err.Error()})
		return
	}
	err, list := form.Sqlite(User, Form)
	if err != nil {
		c.JSON(http.StatusOK, _type.Respond{Code: 0, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, _type.Respond{Code: 1, Message: "加载表单内容列表成功", Data: list})
}
