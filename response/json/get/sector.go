package get

import (
	"BasicOA/serve/sector"
	_type "BasicOA/type"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Sector(c *gin.Context) {
	err, list := sector.List()
	if err != nil {
		c.JSON(http.StatusOK, _type.Respond{Code: 0, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, _type.Respond{Code: 1, Message: "加载部门列表成功", Data: list})
}
