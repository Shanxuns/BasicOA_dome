package task

import (
	"BasicOA/serve/mysql"
	"BasicOA/serve/user"
	_type "BasicOA/type"
	"errors"
)

func List(User _type.User) (error, []interface{}) {
	var (
		Task _type.Task
		err  error
		list []interface{}
	)
	if err, User = user.User(User); err != nil {
		return err, list
	}
	rows, _ := mysql.Db.Query(mysql.TaskList, User.Sector)
	for rows.Next() {
		err := rows.Scan(&Task.Id, &Task.Sector, &Task.Title, &Task.Content, &Task.Start, &Task.End)
		if err != nil {
			return errors.New("获取失败"), nil
		}
		var item = make(map[string]interface{})
		item["id"] = Task.Id
		item["title"] = Task.Title
		item["content"] = Task.Content
		item["start"] = Task.Start
		item["end"] = Task.End
		list = append(list, item)
	}
	return nil, list
}
