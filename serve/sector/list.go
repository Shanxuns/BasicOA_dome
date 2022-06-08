package sector

import (
	"BasicOA/serve/mysql"
	_type "BasicOA/type"
	"errors"
)

func List() (error, []interface{}) {
	var (
		_Sector = new(_type.Sector)
	)
	rows, _ := mysql.Db.Query(mysql.SectorList)
	var list []interface{}
	for rows.Next() {
		err := rows.Scan(&_Sector.Id, &_Sector.Name, &_Sector.Admin, &_Sector.Supervisor)
		if err != nil {
			return errors.New("获取失败"), nil
		}
		var item = make(map[string]interface{})
		item["id"] = _Sector.Id
		item["name"] = _Sector.Name
		item["supervisor"] = _Sector.Supervisor
		list = append(list, item)
	}
	return nil, list
}
