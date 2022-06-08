package post

import (
	"BasicOA/response"
	"BasicOA/serve/regular"
	"BasicOA/serve/task"
	_type "BasicOA/type"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TaskList(c *gin.Context) {
	var (
		User _type.User
		err  error
		list []interface{}
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
	if err, list = task.List(User); err != nil {
		c.JSON(http.StatusOK, _type.Respond{Code: 0, Message: err.Error()})
		return

	}
	c.JSON(http.StatusOK, _type.Respond{Code: 1, Message: "获取成功", Data: list})
}

func Release(c *gin.Context) {
	var (
		User _type.User
		Task _type.Task
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
	if Task.Title, _, err = response.PostForm(c, regular.Title, "title"); err != nil {
		c.JSON(http.StatusOK, _type.Respond{Code: 0, Message: err.Error()})
		return
	}
	if Task.Content, _, err = response.PostForm(c, regular.Content, "content"); err != nil {
		c.JSON(http.StatusOK, _type.Respond{Code: 0, Message: err.Error()})
		return
	}
	if Task.End, _, err = response.PostForm(c, regular.DateTime, "end"); err != nil {
		c.JSON(http.StatusOK, _type.Respond{Code: 0, Message: err.Error()})
		return
	}
	err, Task = task.Release(Task, User)
	if err != nil {
		c.JSON(http.StatusOK, _type.Respond{Code: 0, Message: err.Error()})
		return
	}
	task.SubmitTask(2, "任务通知", Task.Sector, Task)
	c.JSON(http.StatusOK, _type.Respond{Code: 1, Message: "发布任务成功", Data: Task})
}
