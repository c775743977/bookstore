package dao

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"bookstore1.4/model"
	"bookstore1.4/utils"
	"fmt"
	_"strconv"
)

func CreateCart(name string) { //每个用户都应该有自己的购物车，所以每当有新用户创建时就应该给他开启一个购物车数据
	var data bson.D
	res := utils.C_users.FindOne(utils.Ctx, bson.D{{"name", name},})
	err := res.Decode(&data)
	if err != nil {
		fmt.Println("res.Decode(&data) error:", err)
		return
	}
	var cart = bson.D{
		{"num", 0},
		{"amount", 0},
		{"user_id", data[0].Value},
	}
	_, err = utils.C_carts.InsertOne(utils.Ctx, cart)
	if err != nil {
		fmt.Println("C_carts.InsertOne error:", err)
		return
	}
}

func GetCart(userid string) *model.Cart {
	var data bson.D
	id, _ := primitive.ObjectIDFromHex(userid)
	res := utils.C_carts.FindOne(utils.Ctx, bson.D{{"user_id", id},})
	err := res.Decode(&data)
	if err != nil {
		fmt.Println("res.Decode(&data) error:", err)
		return nil
	}
	var cart = model.Cart{
		ID : data[0].Value.(primitive.ObjectID).Hex(),
		Amount : data[2].Value.(float64),
		Num : data[1].Value.(int32),
		UserID : data[3].Value.(primitive.ObjectID).Hex(),
	}
	// utils.DBrr.RoundRobin().Where("user_id = ?", userid).Find(&cart)
	cart.Items = GetItems(cart.ID)
	fmt.Println("cart:", cart)
	return &cart
}

func UpdateCart(userid string) {
	id, _ := primitive.ObjectIDFromHex(userid)
	cart := GetCart(userid)
	var total_num int32
	var total_amount float64
	for _, k := range cart.Items {
		total_num += k.Num
		total_amount += k.Amount
	}
	cart.Num = total_num
	cart.Amount = total_amount
	// utils.WDB.Model(&cart).Select("num", "amount").Where("user_id = ?", userid).Updates(cart)
	var data = bson.D{
		{"$set", bson.D{{"num", cart.Num},}},
		{"$set", bson.D{{"amount", cart.Amount},}},
	}
	_, err := utils.C_carts.UpdateOne(utils.Ctx, bson.D{{"user_id", id},}, data)
	if err != nil {
		fmt.Println("C_carts.UpdateOne error:", err)
		return
	}
}

func CleanCart(userid string) {
	var data bson.D
	id, _ := primitive.ObjectIDFromHex(userid)
	res := utils.C_carts.FindOne(utils.Ctx, bson.D{{"user_id", id},})
	err := res.Decode(&data)
	if err != nil {
		fmt.Println("res.Decode(&data) error:", err)
		return 
	}
	// utils.WDB.Where("cart_id = ?", cart.ID).Delete(&model.CartItem{})
	// utils.WDB.Where("id = ?", cart.ID).Updates(model.Cart{Amount : 0, Num : 0,})
	_, err = utils.C_cartitems.DeleteMany(utils.Ctx, bson.D{{"cart_id", data[0].Value}})
	if err != nil {
		fmt.Println("C_cartitems.DeleteMany error:", err)
		return 
	}
	UpdateCart(userid)
}