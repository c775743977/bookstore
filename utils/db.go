package utils

import (
	_"github.com/mysql"
	"fmt"
	"database/sql"
)

var DB *sql.DB
var err error

func init() {
	DB, err = sql.Open("mysql", "root:Chen@123@/bookstore")
	if err != nil {
		fmt.Println("sql.Open error:", err)
	}
}