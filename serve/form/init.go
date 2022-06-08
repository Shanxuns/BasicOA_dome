package form

import (
	"BasicOA/serve/log"
	"errors"
)

var (
	_Type []string
)

func Run() {
	var appName = "Form"
	log.App(appName, "服务已经启动")
	// 载入文件类型
	_Type = append(_Type, "text")
	_Type = append(_Type, "map")
}

func GetType(x int) (error, string) {
	for i := 0; i < len(_Type); i++ {
		if x == i {
			return nil, _Type[i]
		}
	}
	return errors.New("不支持类型"), ""
}
