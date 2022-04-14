package util

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	// Db 数据库句柄
	Db  *sql.DB
	err error
)

func init() {
	Db, err = sql.Open("mysql", "root:123456@tcp(localhost:3306)/zlibrary")
	if err != nil {
		panic(err.Error())
	}
}
