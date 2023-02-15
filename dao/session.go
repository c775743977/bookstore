package dao

import (
	"bookstore1.4/model"
	"bookstore1.4/utils"
	_"fmt"
	_"strconv"
)

func AddSession(sess *model.Session) {  //将cookie信息存到数据库
	utils.DB.Create(sess)
}

func GetSession(uuid string)  *model.Session { //根据uuid查找登录信息
	var sess model.Session
	utils.DB.Where("id = ?", uuid).Find(&sess)
	return &sess
}

func DelSession(uuid string) {
	utils.DB.Where("id = ?", uuid).Delete(&model.Session{})
}