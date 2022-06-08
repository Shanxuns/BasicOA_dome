package mysql

import (
	"database/sql"
)

var (
	Db    *sql.DB
	errDb interface{}
)
