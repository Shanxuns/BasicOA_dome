package form

import (
	"BasicOA/serve/mysql"
	_type "BasicOA/type"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
)

func Add(FormAddJson _type.FormAddJson) error {
	var (
		err       error
		FormAdmin _type.FormAdmin
		key       = "id"
		value     = "null"
	)
	RowForm := mysql.Db.QueryRow(mysql.FormQueryUser, FormAddJson.Form)
	if err = RowForm.Scan(&FormAdmin.Id, &FormAdmin.Name, &FormAdmin.Admin, &FormAdmin.Path, &FormAdmin.Datetime); err != nil {
		return errors.New("表单不存在")
	}
	for i := 0; i < len(FormAddJson.Data); i++ {
		err, keys := GetType(FormAddJson.Data[i].ID)
		if err != nil {
			return errors.New("表单结构错误")
		}
		key += "," + keys + "_" + strconv.Itoa(i)
		value += ",\"" + FormAddJson.Data[i].Value + "\""
	}
	key += ",datetime"
	value += ",CURRENT_TIMESTAMP"
	db, err := sql.Open("sqlite3", FormAdmin.Path)
	_, err = db.Exec("INSERT INTO `form`(" + key + ") values(" + value + ")")
	if err != nil {
		fmt.Print(err)
		return errors.New("填写失败")
	}
	// 关闭
	err = db.Close()
	if err != nil {
		return errors.New("填写失败")
	}
	return nil
}
