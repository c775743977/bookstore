package dao

import (
	_"go.mongodb.org/mongo-driver/bson"
	"bookstore1.4/model"
	"bookstore1.4/utils"
	"fmt"
	_"strconv"
)

func AddSession(sess *model.Session) {  //将cookie信息存到数据库
	// utils.DB.Create(sess)
	err := utils.RDB.HMSet(utils.Ctx, sess.ID, "user_id", sess.UserID, "user_name", sess.Username).Err()
	if err != nil {
		fmt.Println("AddSession error:", err)
	}
}

func GetSession(uuid string)  *model.Session { //根据uuid查找登录信息
	var sess model.Session
	// utils.DB.Where("id = ?", uuid).Find(&sess)
	res1, err := utils.RDB.HGet(utils.Ctx, uuid, "user_id").Result()
	if err != nil {
		fmt.Println("GetSession error:", err)
		return nil
	} else {
		sess.UserID = res1
	}
	res2, err := utils.RDB.HGet(utils.Ctx, uuid, "user_name").Result()
	if err != nil {
		fmt.Println("GetSession error:", err)
		return nil
	} else {
		sess.Username = res2
	}
	return &sess
}

func DelSession(uuid string) {
	// utils.DB.Where("id = ?", uuid).Delete(&model.Session{})
	err := utils.RDB.HDel(utils.Ctx, uuid, "user_id", "user_name").Err()
	if err != nil {
		fmt.Println("DelSession error:", err)
	}
}