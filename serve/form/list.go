package form

import (
	"BasicOA/serve/mysql"
	"BasicOA/serve/user"
	_type "BasicOA/type"
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

// Db 结构类别
func Db(Form int) (error, []interface{}) {
	var (
		FormType = new(_type.FormType)
		i        = 0
	)
	rows, _ := mysql.Db.Query(mysql.FormTypeQuery, Form)
	var list []interface{}
	for rows.Next() {
		err := rows.Scan(&FormType.Id, &FormType.Form, &FormType.Name, &FormType.Type, &FormType.Extra, &FormType.Datetime)
		if err != nil {
			break
		}
		var item = make(map[string]interface{})
		item["id"] = i
		item["name"] = FormType.Name
		item["type"] = FormType.Type
		item["extra"] = FormType.Extra
		list = append(list, item)
		i++
	}
	if i <= 0 {
		return errors.New("表单不存在"), list
	}
	return nil, list
}

func List(User _type.User) (error, []interface{}) {
	var (
		err       error
		FormAdmin _type.FormAdmin
		list      []interface{}
		i         = 0
	)
	if err, User = user.User(User); err != nil {
		return err, list
	}
	rows, _ := mysql.Db.Query(mysql.FormQueryList, User.Id)
	for rows.Next() {
		if err = rows.Scan(&FormAdmin.Id, &FormAdmin.Name, &FormAdmin.Admin, &FormAdmin.Path, &FormAdmin.Datetime); err != nil {
			break
		}
		var item = make(map[string]interface{})
		item["id"] = FormAdmin.Id
		item["name"] = FormAdmin.Name
		item["datetime"] = FormAdmin.Datetime
		list = append(list, item)
		i++
	}
	if i <= 0 {
		return errors.New("未创建表单"), list
	}
	return nil, list
}

func Sqlite(User _type.User, Form int) (error, []interface{}) {
	var (
		err       error
		FormAdmin _type.FormAdmin
		list      []interface{}
		i         = 0
		FormType  = new(_type.FormType)
		x         = 2
		_dBType   []string
	)
	if err, User = user.User(User); err != nil {
		return err, list
	}
	if err, FormAdmin = Admin(User, Form); err != nil {
		return err, list
	}
	_, err = os.OpenFile(FormAdmin.Path, os.O_APPEND, 0666)
	if err != nil {
		return errors.New("未创建表单结构"), list
	}

	rows, _ := mysql.Db.Query(mysql.FormTypeQuery, Form)
	_dBType = append(_dBType, "序号")
	for rows.Next() {
		err := rows.Scan(&FormType.Id, &FormType.Form, &FormType.Name, &FormType.Type, &FormType.Extra, &FormType.Datetime)
		if err != nil {
			break
		}
		_dBType = append(_dBType, FormType.Name)
		x++
	}
	_dBType = append(_dBType, "填写时间")
	if x <= 2 {
		return errors.New("表单不存在"), list
	}
	db, err := sql.Open("sqlite3", FormAdmin.Path)
	rowsSqlite, _ := db.Query(mysql.FormSqlite)
	for rowsSqlite.Next() {
		value := make([]interface{}, x)
		data := make([][]byte, x) //数据库中的NULL值可以扫描到字节中
		for i := range value {
			value[i] = &data[i]
		}
		if err = rowsSqlite.Scan(value...); err != nil {
			break
		}
		var item = make(map[string]interface{})
		for k, col := range data {
			item[_dBType[k]] = string(col)
		}
		list = append(list, item)
		i++
	}
	if i <= 0 {
		return errors.New("没有填写数据"), list
	}
	// 关闭
	err = db.Close()
	if err != nil {
		return errors.New("创建表单数据库关闭失败"), list
	}
	return nil, list
}
