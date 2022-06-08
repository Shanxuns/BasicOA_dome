package user

import (
	"BasicOA/serve/md5"
	"BasicOA/serve/mysql"
	_type "BasicOA/type"
	"errors"
)

func Activate(Sector _type.Sector, User _type.User, Activate string) (error, _type.Sector) {
	// Mysql查找部门
	RowSector := mysql.Db.QueryRow(mysql.Sector, Sector.Id)
	if err := RowSector.Scan(&Sector.Id, &Sector.Name, &Sector.Admin, &Sector.Supervisor); err != nil {
		return errors.New("无当前部门"), Sector
	}
	// Mysql查找账户
	RowUser := mysql.Db.QueryRow(mysql.SearchEmail, User.Email)
	if err := RowUser.Scan(&User.Id, &User.Email, &User.Fname, &User.Lname, &User.Phone, &User.Password, &User.Confuse, &User.Avatar, &User.License, &User.Datetime); err != nil {
		return errors.New("账号未注册"), Sector
	}
	// 验证
	if User.License == 1 {
		return errors.New("账号已激活"), Sector
	}
	md5.SetHash("Activation")
	if md5.Hash(User.Confuse) != Activate {
		return errors.New("激活码错误"), Sector
	}
	// 更新激活状态
	if _, err := mysql.Db.Exec(mysql.Activate, 1, User.Id); err != nil {
		return errors.New("激活失败请联系管理员"), Sector
	}
	if _, err := mysql.Db.Exec(mysql.Superior, &User.Id, &Sector.Id); err != nil {
		return errors.New("申请加入部门失败"), Sector
	}
	return nil, Sector
}
