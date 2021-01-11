package repositories

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)


func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:Blacklotus@tcp(127.0.0.1:3306)/schoolDB")
	if err != nil {
		panic(err.Error())
	}
	return db

}
