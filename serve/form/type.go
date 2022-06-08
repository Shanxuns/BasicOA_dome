package form

import (
	"BasicOA/serve/mysql"
	"BasicOA/serve/user"
	_type "BasicOA/type"
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"strconv"
)

func Type(User _type.User, formTypeJson _type.FormTypeJson) error {
	var (
		err       error
		FormAdmin _type.FormAdmin
		table     = "CREATE TABLE IF NOT EXISTS `form` (id INTEGER PRIMARY KEY AUTOINCREMENT"
	)
	if err, User = user.User(User); err != nil {
		return err
	}
	if err, FormAdmin = Admin(User, formTypeJson.Form); err != nil {
		return err
	}
	_, err = os.OpenFile(FormAdmin.Path, os.O_APPEND, 0666)
	if err == nil {
		return errors.New("已经创建表单结构")
	}
	for i := 0; i < len(formTypeJson.Type); i++ {
		var name = formTypeJson.Type[i].Name
		var _Type = formTypeJson.Type[i].Type
		var extra = formTypeJson.Type[i].Extra
		err, key := GetType(_Type)
		if err != nil {
			return errors.New("表单结构错误")
		}
		if _, err = mysql.Db.Exec(mysql.FormType, FormAdmin.Id, name, _Type, extra); err != nil {
			return errors.New("创建表单结构失败")
		}
		table += "," + key + "_" + strconv.Itoa(i) + " VARCHAR(128) NULL"
	}
	table += ",datetime DATETIME DEFAULT CURRENT_TIMESTAMP)"
	db, err := sql.Open("sqlite3", FormAdmin.Path)
	_, err = db.Exec(table)
	if err != nil {
		return err
	}
	// 关闭
	err = db.Close()
	if err != nil {
		return errors.New("创建表单数据库关闭失败")
	}
	return nil
}
