package user

import (
	"BasicOA/serve/mysql"
	_type "BasicOA/type"
	"errors"
)

func User(User _type.User) (error, _type.User) {
	// Mysql查找账户

	var Superior = new(_type.Superior)
	RowUser := mysql.Db.QueryRow(mysql.User, User.Email, User.Password)
	if err := RowUser.Scan(&User.Id, &User.Email, &User.Fname, &User.Lname, &User.Phone, &User.Password, &User.Confuse, &User.Avatar, &User.License, &User.Datetime); err != nil {
		return errors.New("账号或密码错误"), User
	}
	if User.License == 0 {
		return errors.New("账号未激活"), User
	}
	RowSuperior := mysql.Db.QueryRow(mysql.SearchSuperior, User.Id)
	if err := RowSuperior.Scan(&Superior.Id, &Superior.User, &Superior.Sector, &Superior.License, &Superior.Datetime); err != nil {
		return errors.New("账号未激活"), User
	}
	if Superior.License == 0 {
		return errors.New("部门未同意申请"), User
	}
	User.Sector = Superior.Sector
	User.Password = ""
	User.Confuse = ""
	User.License = 0
	return nil, User
}
