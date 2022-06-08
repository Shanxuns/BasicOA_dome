package gin

import (
	"BasicOA/response/json/get"
	"BasicOA/response/json/post"
	"BasicOA/serve/task"
)

func runRoute() {
	// V1 版本
	Api := Gin.Group("/v1/api")
	{
		// json
		Json := Api.Group("/json")
		// 登录
		Json.POST("/Login", post.Login)
		// 注册-
		Json.POST("/Register", post.Register)
		// 激活
		Json.POST("/Activate", post.Activate)
		// 管理员操作用户
		Json.POST("/Operate", post.Operate)
		// 用户信息
		Json.GET("/User", get.User)
		// 部门类别
		Json.GET("/Sector", get.Sector)
		// 任务系统WebSocketTask
		Json.GET("/WebSocketTask", task.WebSocketTask)
		// 发布任务
		Json.POST("/Release", post.Release)
		// 任务列表
		Json.POST("/TaskList", post.TaskList)
		// 新建表单
		Json.POST("/FormNew", post.FormNew)
		// 新建表单结构
		Json.POST("/FormTypeNew", post.FormTypeNew)
		// 表单数据结构类别
		Json.GET("/FormTypeList/:id", get.TypeList)
		// 填写表单
		Json.POST("/FormAdd", post.FormAdd)
		// 表单信息
		Json.GET("/FormTitle/:id", get.Title)
		// 表单列表
		Json.GET("/FormList", get.List)
		// 表单内容列表
		Json.GET("/FormContent/:id", get.Sqlite)
	}
}
