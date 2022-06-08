package _type

import "BasicOA/serve/log"

var (
	// FileType 文件类型
	FileType []string
	// PermitType 文件操作数
	PermitType []string
)

func Run() {
	var appName = "File"
	log.App(appName, "服务已经启动")
	// 载入文件类型
	FileType = append(FileType, ".bin")
	FileType = append(FileType, "_folder")
	FileType = append(FileType, ".jpg")
	// 载入文件操作数
	PermitType = append(PermitType, "移动")
	PermitType = append(PermitType, "复制")
	PermitType = append(PermitType, "新建")
	PermitType = append(PermitType, "分享")
	PermitType = append(PermitType, "删除")
	PermitType = append(PermitType, "上传")
	PermitType = append(PermitType, "下载")
}

// Respond 一般响应
type Respond struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
	Task    int         `json:"Task,omitempty"`
	Data    interface{} `json:"data"`
}

// User 用户信息
type User struct {
	Id       int    `json:"id,omitempty" db:"id"`
	Email    string `json:"email,omitempty" db:"email"`
	Fname    string `json:"fname,omitempty" db:"fname"`
	Lname    string `json:"lname,omitempty" db:"lname"`
	Phone    string `json:"phone,omitempty" db:"phone"`
	Password string `json:"password,omitempty" db:"password"`
	Confuse  string `json:"confuse,omitempty" db:"confuse"`
	Avatar   string `json:"avatar,omitempty" db:"avatar"`
	License  int    `json:"license,omitempty" db:"license"`
	Datetime string `json:"datetime,omitempty" db:"datetime"`
	Sector   int    `json:"sector,omitempty"`
}

// Sector  部门信息
type Sector struct {
	Id         int    `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	Admin      int    `json:"admin" db:"admin"`
	Supervisor int    `json:"supervisor" db:"supervisor"`
}

// Superior 用户上级部门信息
type Superior struct {
	Id       int    `json:"id" db:"id"`
	User     int    `json:"user" db:"user"`
	Sector   int    `json:"sector" db:"sector"`
	License  int    `json:"license" db:"license"`
	Datetime string `json:"datetime" db:"datetime"`
}

// File 文件信息
type File struct {
	Id         int    `json:"id" db:"id"`
	Folder     int    `json:"folder" db:"folder"`
	Belong     int    `json:"belong" db:"belong"`
	BelongType int    `json:"belong_type" db:"belong_type"`
	Name       string `json:"name" db:"name"`
	Type       int    `json:"type" db:"type"`
	Path       string `json:"path" db:"path"`
	Datetime   string `json:"datetime" db:"datetime"`
}

// Task 任务信息
type Task struct {
	Id      int    `json:"id,omitempty" db:"id"`
	Sector  int    `json:"sector,omitempty" db:"sector"`
	Title   string `json:"title,omitempty" db:"title"`
	Content string `json:"content,omitempty" db:"content"`
	Start   string `json:"start,omitempty" db:"start"`
	End     string `json:"end,omitempty" db:"end"`
}

// FormAdmin 表单管理员
type FormAdmin struct {
	Id       int    `json:"id,omitempty" db:"id"`
	Name     string `json:"name,omitempty" db:"name"`
	Admin    int    `json:"admin,omitempty" db:"admin"`
	Path     string `json:"path,omitempty" db:"path"`
	Datetime string `json:"datetime,omitempty" db:"datetime"`
}

// FormType 表单数据结构
type FormType struct {
	Id       int    `json:"id,omitempty" db:"id"`
	Form     int    `json:"form,omitempty" db:"form"`
	Name     string `json:"name,omitempty" db:"name"`
	Type     int    `json:"type,omitempty" db:"type"`
	Extra    string `json:"extra,omitempty" db:"extra"`
	Datetime string `json:"datetime,omitempty" db:"datetime"`
}

// FormTypeJson 表单结构Json
type FormTypeJson struct {
	Form int `json:"form"`
	Type []struct {
		Name  string `json:"name"`
		Type  int    `json:"type"`
		Extra string `json:"extra"`
	} `json:"type"`
}

// FormAddJson  填写表单
type FormAddJson struct {
	Form int `json:"form"`
	Data []struct {
		ID    int    `json:"id"`
		Value string `json:"value"`
	} `json:"data"`
}
