package dao

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"bookstore1.4/model"
	"bookstore1.4/utils"
	_"strconv"
	"fmt"
)

func CheckUserName(name string) bool {
	res := utils.C_users.FindOne(utils.Ctx, bson.D{{"name", name},})
	err := res.Decode(&bson.D{})
	if err != nil {
		fmt.Println("res.Decode(&user) error:", err)
		return true
	} else {
		return false
	}
}

// func CheckUserName(name string) bool { //仅验证用户名，用于注册时，如果想要用户名可以重复，那就以userid作为唯一标识符
// 	var user model.User //注意如果用var user *model.User定义，在Find(user)时无法传入数据(要么用var user model.User要么用user:=&model.User{})
// 	utils.DBrr.RoundRobin().Where("name = ?", name).Find(&user)
// 	if user.ID != 0 {
// 		return false
// 	} else {
// 		return true
// 	}
// }

func CheckUserNameAndPassword(name string, passwd string) string { //验证用户名和密码，用于登录
	var user bson.D
	res := utils.C_users.FindOne(utils.Ctx, bson.D{{"name",name},{"password",passwd},})
	err := res.Decode(&user)
	if err != nil {
		fmt.Println("res.Decode(&user) error:", err)
		return ""
	}
	// user := &model.User{}
	// utils.DBrr.RoundRobin().Select("id", "password").Where("name = ?", name).Find(user)
	uuid := utils.CreateUUID()
	sess := &model.Session{
		ID : uuid,
		Username : name,
		UserID : user[0].Value.(primitive.ObjectID).Hex(),
	}
	AddSession(sess)
	return uuid
}

func UserRegist(user *model.User) {  //注册新用户
	var data = bson.D{
		{"name", user.Name},
		{"password", user.Password},
		{"email", user.Email},
		{"privilege", "N"},
	}
	_, err := utils.C_users.InsertOne(utils.Ctx, data)
	if err != nil {
		fmt.Println("C_users.InsertOne error:", err)
		return
	}
}

func GetAllUsers() []*model.User {
	var us []*model.User
	var data bson.D
	cursor, err := utils.C_users.Find(utils.Ctx, bson.D{})
	if err != nil {
		fmt.Println("utils.C_users error:", err)
		return nil
	}
	for cursor.Next(utils.Ctx) {
		cursor.Decode(&data)
		user := &model.User{
			ID : data[0].Value.(primitive.ObjectID).Hex(),
			Name : data[1].Value.(string),
			Password : data[2].Value.(string),
			Email : data[3].Value.(string),
			Privilege : data[4].Value.(string),
		}
		us = append(us, user)
	}
	return us
}

func AlterUser(user *model.User) { //更改用户信息
	// utils.WDB.Where("id = ?", user.ID).Save(user)
	fmt.Println("user:", user)
	uid, _ := primitive.ObjectIDFromHex(user.ID)
	_, err := utils.C_users.UpdateOne(utils.Ctx, bson.D{{"_id", uid},}, bson.D{
		{"$set", bson.D{{"name", user.Name},}},
		{"$set", bson.D{{"password", user.Password},}},
		{"$set", bson.D{{"email", user.Email},}},
		{"$set", bson.D{{"privilege", user.Privilege},}},
	})
	fmt.Println("ok")
	if err != nil {
		fmt.Println("C_users.UpdateOne error:", err)
		return
	}
}

func DelUser(id string) bool {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println("primitive.ObjectIDFromHex error:", err)
		return false
	}
	var data bson.D
	var order_id []string
	res := utils.C_carts.FindOne(utils.Ctx, bson.D{{"_id", ID},})
	err = res.Decode(&data)
	if err != nil {
		fmt.Println("res.Decode(&data) error1:", err)
		return false
	}
	cart_id := data[0].Value
	utils.C_carts.DeleteOne(utils.Ctx, bson.D{{"_id", cart_id},})
	utils.C_carts.DeleteMany(utils.Ctx, bson.D{{"cart_id", cart_id},})
	cursor, err := utils.C_orders.Find(utils.Ctx, bson.D{{"user_id", ID},})
	if err != nil {
		fmt.Println("C_orders.Find error:", err)
		return false
	}
	for cursor.Next(utils.Ctx) {
		err = cursor.Decode(&data)
		if err != nil {
			fmt.Println("res.Decode(&data) error2:", err)
		}
		order_id = append(order_id, data[0].Value.(string))
	}
	utils.C_orders.DeleteMany(utils.Ctx, bson.D{{"user_id", ID},})
	for _, k := range order_id {
		utils.C_orderitems.DeleteMany(utils.Ctx, bson.D{{"order_id", k},})
	}
	return true
}

// func DelUser(id string) bool  {
// 	var cart model.Cart
// 	utils.DBrr.RoundRobin().Where("user_id = ?", userid).Find(&cart)
// 	var orders []*model.Order
// 	result := utils.DBrr.RoundRobin().Where("user_id = ?", userid).Find(&orders)
// 	for _, k := range orders {
// 		if k.Status != 4 {
// 			return false
// 		}
// 	}
// 	utils.WDB.Where("cart_id = ?", cart.ID).Delete(&model.CartItem{})
// 	utils.WDB.Where("user_id = ?", userid).Delete(&model.Cart{})
// 	if result.RowsAffected != 0 {
// 		utils.WDB.Where("order_id = ?", orders[0].ID).Delete(&model.OrderItem{})
// 		utils.WDB.Where("user_id = ?", userid).Delete(&model.Order{})
// 	}
// 	utils.WDB.Where("id = ?", userid).Delete(&model.User{})
// 	return true
// }

func GetUserByID(userid string) *model.User {
	var data bson.D
	uid, _ := primitive.ObjectIDFromHex(userid)
	res := utils.C_users.FindOne(utils.Ctx, bson.D{{"_id", uid},})
	err := res.Decode(&data)
	if err != nil {
		fmt.Println("C_users.FindOne error:", err)
		return nil
	}
	user := &model.User{
		ID : data[0].Value.(primitive.ObjectID).Hex(),
		Name : data[1].Value.(string),
		Password : data[2].Value.(string),
		Email : data[3].Value.(string),
		Privilege : data[4].Value.(string),
	}
	return user
}