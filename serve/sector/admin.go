package sector

import (
	"BasicOA/serve/mysql"
	"BasicOA/serve/user"
	_type "BasicOA/type"
	"errors"
)

func Admin(User _type.User) (error, _type.User) {
	var (
		err    error
		Sector _type.Sector
	)
	if err, User = user.User(User); err != nil {
		return err, User
	}
	RowSector := mysql.Db.QueryRow(mysql.AdminSector, User.Id)
	if err := RowSector.Scan(&Sector.Id, &Sector.Name, &Sector.Admin, &Sector.Supervisor); err != nil {
		return errors.New("无权限"), User
	}
	return nil, User
}
