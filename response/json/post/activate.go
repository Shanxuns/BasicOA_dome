package post

import (
	"BasicOA/response"
	"BasicOA/serve/regular"
	"BasicOA/serve/task"
	"BasicOA/serve/user"
	_type "BasicOA/type"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Activate(c *gin.Context) {
	// 创建User结构体
	var (
		User      _type.User
		Sector    _type.Sector
		err       error
		_Activate string
	)
	// 写入并验证参数
	if User.Email, _, err = response.PostForm(c, regular.Email, "email"); err != nil {
		c.JSON(http.StatusOK, _type.Respond{Code: 0, Message: err.Error()})
		return
	}
	if _, Sector.Id, err = response.PostForm(c, regular.Id, "sector"); err != nil {
		c.JSON(http.StatusOK, _type.Respond{Code: 0, Message: err.Error()})
		return
	}
	if _Activate, _, err = response.PostForm(c, regular.Md5, "activate"); err != nil {
		c.JSON(http.StatusOK, _type.Respond{Code: 0, Message: err.Error()})
		return
	}
	if err, Sector = user.Activate(Sector, User, _Activate); err != nil {
		c.JSON(http.StatusOK, _type.Respond{Code: 0, Message: err.Error()})
		return
	}
	task.SubmitTask(3, "申请通知", Sector.Admin, User)
	c.JSON(http.StatusOK, _type.Respond{Code: 1, Message: "激活成功"})
	return
}
