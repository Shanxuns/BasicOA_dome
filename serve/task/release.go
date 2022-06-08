package task

import (
	"BasicOA/serve/mysql"
	"BasicOA/serve/sector"
	_type "BasicOA/type"
	"errors"
	"time"
)

func Release(Task _type.Task, User _type.User) (error, _type.Task) {
	var Sector _type.Sector
	t1, err := time.Parse("2006-01-02 15:04:05", Task.End)
	if err == nil && t1.Before(time.Now()) {
		return errors.New("任务结束时间小于任务开始时间"), Task
	}
	if err, _ = sector.Admin(User); err != nil {
		return err, Task
	}
	Task.Sector = Sector.Id
	// 任务信息写入数据库
	task, err := mysql.Db.Exec(mysql.TaskRelease, Task.Sector, Task.Title, Task.Content, Task.End)
	if err != nil {
		return errors.New("发布任务失败"), Task
	}
	id, err := task.LastInsertId()
	if err != nil {
		return errors.New("推送任务失败"), Task
	}
	Task.Id = int(id)
	Task.Content = ""
	return nil, Task
}
