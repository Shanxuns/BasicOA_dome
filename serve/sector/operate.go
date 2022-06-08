package sector

import (
	"BasicOA/serve/mysql"
	_type "BasicOA/type"
	"errors"
)

func Operate(User _type.User, Superior _type.Superior) error {
	err, User := Admin(User)
	if err != nil {
		return err
	}
	if User.Id == Superior.User {
		return errors.New("不能为自己修改激活状态")
	}
	RowUser := mysql.Db.QueryRow(mysql.SearchId, Superior.User)
	if err := RowUser.Scan(&User.Id, &User.Email, &User.Fname, &User.Lname, &User.Phone, &User.Password, &User.Confuse, &User.Avatar, &User.License, &User.Datetime); err != nil {
		return errors.New("账号未注册")
	}
	if User.License == 0 {
		return errors.New("账号未激活")
	}
	if _, err := mysql.Db.Exec(mysql.Operate, Superior.License, Superior.User, User.Sector); err != nil {
		return errors.New("操作失败")
	}
	return nil
}
