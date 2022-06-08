package user

import (
	"BasicOA/serve/md5"
	"BasicOA/serve/mysql"
	_type "BasicOA/type"
	"errors"
)

func Login(User _type.User) (error, _type.User) {
	Password := User.Password
	// Mysql查找账户
	RowUser := mysql.Db.QueryRow(mysql.SearchEmail, User.Email)
	if err := RowUser.Scan(&User.Id, &User.Email, &User.Fname, &User.Lname, &User.Phone, &User.Password, &User.Confuse, &User.Avatar, &User.License, &User.Datetime); err != nil {
		return errors.New("账号未注册"), User
	}
	md5.SetHash(User.Confuse)
	if md5.Hash(Password) != User.Password {
		return errors.New("账号或密码错误"), User
	}
	if User.License == 0 {
		return errors.New("账号未激活"), User
	}
	User.Confuse = ""
	User.License = 0
	return nil, User
}
 