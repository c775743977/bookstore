package utils

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"fmt"
)

var DB *gorm.DB
var err error

func init() {
	DB, err = gorm.Open(mysql.Open("root:Chen@123@/bookstore"), &gorm.Config{})
	if err != nil {
		fmt.Println("connect to db error:", err)
	}
}