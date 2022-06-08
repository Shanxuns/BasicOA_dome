package form

import (
	"BasicOA/serve/mysql"
	_type "BasicOA/type"
	"errors"
)

func Title(id int) (error, _type.FormAdmin) {
	var (
		err       error
		FormAdmin _type.FormAdmin
	)
	RowForm := mysql.Db.QueryRow(mysql.FormQueryUser, id)
	if err = RowForm.Scan(&FormAdmin.Id, &FormAdmin.Name, &FormAdmin.Admin, &FormAdmin.Path, &FormAdmin.Datetime); err != nil {
		return errors.New("未创建表单"), FormAdmin
	}
	return nil, FormAdmin
}
