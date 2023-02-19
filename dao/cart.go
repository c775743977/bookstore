package dao

import (
	"bookstore1.4/model"
	"bookstore1.4/utils"
	_"fmt"
	_"strconv"
)

func CreateCart(name string) { //每个用户都应该有自己的购物车，所以每当有新用户创建时就应该给他开启一个购物车数据
	var id int
	utils.DBrr.RoundRobin().Model(&model.User{}).Select("id").Where("name = ?", name).Find(&id)
	var cart = model.Cart{
		Num : 0,
		Amount : 0,
		UserID : id,
	}
	utils.WDB.Create(&cart)
}

func GetCart(userid int) *model.Cart {
	var cart model.Cart
	utils.DBrr.RoundRobin().Where("user_id = ?", userid).Find(&cart)
	cart.Items = GetItems(cart.ID)
	return &cart
}

func UpdateCart(userid int) {
	cart := GetCart(userid)
	var total_num int
	var total_amount float64
	for _, k := range cart.Items {
		total_num += k.Num
		total_amount += k.Amount
	}
	cart.Num = total_num
	cart.Amount = total_amount
	utils.WDB.Model(&cart).Select("num", "amount").Where("user_id = ?", userid).Updates(cart)
}

func CleanCart(userid int) {
	cart := GetCart(userid)
	utils.WDB.Where("cart_id = ?", cart.ID).Delete(&model.CartItem{})
	utils.WDB.Where("id = ?", cart.ID).Updates(model.Cart{Amount : 0, Num : 0,})
	UpdateCart(userid)
}