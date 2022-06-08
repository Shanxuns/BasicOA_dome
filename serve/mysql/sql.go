package mysql

var (
	// Register 添加用户
	Register = "INSERT INTO user(`id`, `email`, `fname`, `lname`, `phone`, `password`, `confuse`, `avatar`, `license`, `datetime`) VALUES (null, ?, ?, ?, ?, ?, ?, ?, 0, CURRENT_TIMESTAMP)"
	// SearchEmail 电子邮箱搜索用户
	SearchEmail = "select * from user where email=?"
	// SearchId 电子邮箱搜索用户
	SearchId = "select * from user where id=?"
	// Activate 激活用户
	Activate = "UPDATE user set license=? where id=?"
	// Operate 管理员部门用户操作
	Operate = "UPDATE superior set license=? where user=? and sector=?"
	// User 验证用户
	User = "select * from user where email=? and password=?"
	// Superior  申请加入部门
	Superior = "INSERT INTO superior(`id`, `user`, `sector`, `license`, `datetime`) VALUES (NULL, ?, ?, 0, CURRENT_TIMESTAMP)"
	// SearchSuperior 用户部门
	SearchSuperior = "select * from superior where user=?"
	// Sector 查询部门
	Sector = "select * from sector where id=?"
	// AdminSector 管理员部门
	AdminSector = "select * from sector where admin=?"
	// SectorList 部门列表
	SectorList = "select * from sector order by id desc"
	// TaskRelease 发布任务
	TaskRelease = "INSERT INTO `task`(`id`, `sector`, `title`, `content`, `start`, `end`) VALUES (NULL, ?, ?, ?, CURRENT_TIMESTAMP, ?)"
	// TaskList 任务列表
	TaskList = "select * from `task` where sector=? or sector=0 order by id desc"
	// FormAdmin 新建表单
	FormAdmin = "INSERT INTO `form_admin`(`id`, `name`, `admin`, `path`, `datetime`) VALUES (NULL, ?, ?, ?, CURRENT_TIMESTAMP)"
	// FormQuery 表单查询
	FormQuery = "select * from `form_admin` where id=? and admin=?"
	// FormQueryUser 表单查询
	FormQueryUser = "select * from `form_admin` where id=?"
	// FormQueryList 表单列表
	FormQueryList = "select * from `form_admin` where admin=?"
	// FormType 新建表单结构
	FormType = "INSERT INTO `form_db`(`id`, `form`, `name`, `type`, `extra`, `datetime`) VALUES (NULL, ?, ?, ?, ?, CURRENT_TIMESTAMP)"
	// FormTypeQuery 表单查询
	FormTypeQuery = "select * from `form_db` where form=?"
	// FormSqlite sqlite 数据库 查讯
	FormSqlite = "select * from `form`"
)
