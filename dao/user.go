package dao

import (
	"Hello_golang/bookstore/model"
	"Hello_golang/bookstore/utils"
	"fmt"
)

func CheckUserName(name string) bool { //仅验证用户名，用于注册时，如果想要用户名可以重复，那就以userid作为唯一标识符
	sqlstr := "select password from bookstore.users where name=?"
	row := utils.DB.QueryRow(sqlstr, name)
	var res string
	_ = row.Scan(&res)
	if res != "" {
		return false
	} else {
		return true
	}
}

func CheckUserNameAndPassword(name string, passwd string) string { //验证用户名和密码，用于登录
	sqlstr := "select id,password from bookstore.users where name=?"
	row := utils.DB.QueryRow(sqlstr, name)
	var res string
	var id int
	err := row.Scan(&id, &res)
	if err != nil {
		fmt.Println("row.Scan error:", err)
		return ""
	}
	if res == passwd {
		uuid := utils.CreateUUID()
		sess := &model.Session{
			ID : uuid,
			Username : name,
			UserID : id,
		}
		AddSession(sess)
		return uuid
	} else {
		return ""
	}
}

func UserRegister(name string, passwd string, email string) {  //注册新用户
	sqlstr := "insert into bookstore.users(name, password, email) values(?,?,?)"
	_, err := utils.DB.Exec(sqlstr, name, passwd, email)
	if err != nil {
		fmt.Println("utils.DB.Exec error:", err)
	}
}