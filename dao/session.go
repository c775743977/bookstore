package dao

import (
	"bookstore1.2/model"
	"bookstore1.2/utils"
	"fmt"
	_"strconv"
)

func AddSession(sess *model.Session) {
	sqlstr := "insert into bookstore.session value(?,?,?)"
	_ ,err := utils.DB.Exec(sqlstr, sess.ID, sess.Username, sess.UserID)
	if err != nil {
		fmt.Println("AddSession utils.DB.Exec error:", err)
		return
	}
}

func DelSession(sessionID string) {
	sqlstr := "delete from bookstore.session where id = ?"
	_, err := utils.DB.Exec(sqlstr, sessionID)
	if err != nil {
		fmt.Println("DelSession utils.DB.Exec error:", err)
		return
	}
}


func GetSession(sessID string) (username string, userid int) {
	sqlstr := "select username, userid from session where id=?"
	row := utils.DB.QueryRow(sqlstr, sessID)
	sess := &model.Session{}
	sess.ID = sessID
	err := row.Scan(&sess.Username, &sess.UserID)
	if err != nil {
		fmt.Println("GetSession row.Scan error:", err)
		return "", 0
	}
	return sess.Username, sess.UserID
}
