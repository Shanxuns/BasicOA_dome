package form

import (
	"BasicOA/serve/md5"
	"BasicOA/serve/mysql"
	"BasicOA/serve/user"
	_type "BasicOA/type"
	"errors"
	"strconv"
	"time"
)

func New(User _type.User, name string) error {
	var (
		err  error
		path string
	)
	if err, User = user.User(User); err != nil {
		return err
	}
	md5.SetHash(User.Email)
	path = "./sqlite/" + md5.Hash(strconv.FormatInt(time.Now().UnixNano(), 10)) + ".db"
	if err != nil {
		return errors.New("创建失败")
	}
	if _, err = mysql.Db.Exec(mysql.FormAdmin, name, User.Id, path); err != nil {
		return errors.New("创建失败")
	}
	return nil
}
