package post

import (
	"BasicOA/response"
	"BasicOA/serve/form"
	"BasicOA/serve/regular"
	_type "BasicOA/type"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FormNew(c *gin.Context) {
	var (
		User _type.User
		err  error
		name string
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
	if name, _, err = response.PostForm(c, regular.Title, "name"); err != nil {
		c.JSON(http.StatusOK, _type.Respond{Code: 0, Message: err.Error()})
		return
	}
	if err = form.New(User, name); err != nil {
		c.JSON(http.StatusOK, _type.Respond{Code: 0, Message: err.Error()})
		return
	}
	// 响应结果
	c.JSON(http.StatusOK, _type.Respond{Code: 1, Message: "新建成功"})
}

func FormTypeNew(c *gin.Context) {
	var (
		User         _type.User
		FormTypeJson _type.FormTypeJson
		err          error
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
	err = c.BindJSON(&FormTypeJson)
	if err != nil {
		c.JSON(http.StatusOK, _type.Respond{Code: 0, Message: "请求参数错误"})
		return
	}
	if err = form.Type(User, FormTypeJson); err != nil {
		c.JSON(http.StatusOK, _type.Respond{Code: 0, Message: err.Error()})
		return
	}
	// 响应结果
	c.JSON(http.StatusOK, _type.Respond{Code: 1, Message: "创建成功"})
}

func FormAdd(c *gin.Context) {
	var (
		FormAddJson _type.FormAddJson
		err         error
	)
	err = c.BindJSON(&FormAddJson)
	if err != nil {
		c.JSON(http.StatusOK, _type.Respond{Code: 0, Message: "请求参数错误"})
		return
	}
	if err = form.Add(FormAddJson); err != nil {
		c.JSON(http.StatusOK, _type.Respond{Code: 0, Message: err.Error()})
		return
	}
	// 响应结果
	c.JSON(http.StatusOK, _type.Respond{Code: 1, Message: "填写成功"})
}
