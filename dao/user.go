package dao

import (
	"bookstore1.4/model"
	"bookstore1.4/utils"
	"strconv"
	_"fmt"
)

func CheckUserName(name string) bool { //仅验证用户名，用于注册时，如果想要用户名可以重复，那就以userid作为唯一标识符
	var user model.User //注意如果用var user *model.User定义，在Find(user)时无法传入数据(要么用var user model.User要么用user:=&model.User{})
	utils.DB.Where("name = ?", name).Find(&user)
	if user.ID != 0 {
		return false
	} else {
		return true
	}
}

func CheckUserNameAndPassword(name string, passwd string) string { //验证用户名和密码，用于登录
	user := &model.User{}
	utils.DB.Select("id", "password").Where("name = ?", name).Find(user)
	if user.Password == passwd {
		uuid := utils.CreateUUID()
		sess := &model.Session{
			ID : uuid,
			Username : name,
			UserID : user.ID,
		}
		AddSession(sess)
		return uuid
	} else {
		return ""
	}
}

func UserRegist(user *model.User) {  //注册新用户
	utils.DB.Create(user)
}

func GetAllUsers() []*model.User {
	var us []*model.User
	utils.DB.Find(&us)
	return us
}

func AlterUser(user *model.User) { //更改用户信息
	utils.DB.Where("id = ?", user.ID).Save(user)
}

func DelUser(id string) bool  {
	userid, _ := strconv.ParseInt(id, 10, 0)
	var cart model.Cart
	utils.DB.Where("user_id = ?", userid).Find(&cart)
	var orders []*model.Order
	result := utils.DB.Where("user_id = ?", userid).Find(&orders)
	for _, k := range orders {
		if k.Status != 4 {
			return false
		}
	}
	utils.DB.Where("cart_id = ?", cart.ID).Delete(&model.CartItem{})
	utils.DB.Where("user_id = ?", userid).Delete(&model.Cart{})
	if result.RowsAffected != 0 {
		utils.DB.Where("order_id = ?", orders[0].ID).Delete(&model.OrderItem{})
		utils.DB.Where("user_id = ?", userid).Delete(&model.Order{})
	}
	utils.DB.Where("id = ?", userid).Delete(&model.User{})
	return true
}

func GetUserByID(userid int) *model.User {
	var user model.User
	utils.DB.Where("id = ?", userid).Find(&user)
	return &user
}