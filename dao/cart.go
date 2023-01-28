package dao

import (
	"Hello_golang/bookstore/model"
	"Hello_golang/bookstore/utils"
	"fmt"
	"strconv"
)

func CreateCart(username string) {  //每个用户都应该有自己的购物车，所以每当有新用户创建时就应该给他开启一个购物车数据
	sqlstr := "select id from bookstore.users where name = ?"
	row := utils.DB.QueryRow(sqlstr, username)
	var temp string
	err := row.Scan(&temp)
	if err != nil {
		fmt.Println("CreateCart row.Scan error:", err)
		return
	}
	userid, _ := strconv.ParseInt(temp, 0, 0)
	sqlstr = "insert into bookstore.carts(num,amount,userid) values(0,0,?)"
	_, err = utils.DB.Exec(sqlstr, userid)
	if err != nil {
		fmt.Println("CreateCart utils.DB.Exec error:", err)
		return
	}
}

func GetCart(userid int) *model.Cart {
	cart := &model.Cart{}
	sqlstr := "select * from bookstore.carts where userid = ?"
	row := utils.DB.QueryRow(sqlstr, userid)
	err := row.Scan(&cart.ID, &cart.Num, &cart.Amount, &cart.UserID)
	cart.Items = GetItems(cart.ID)
	if err != nil {
		fmt.Println("GetCart row.Scan error:", err)
		return nil
	}
	return cart
}

func UpdateCart(userid int) {
	cart := GetCart(userid)
	var total_num int
	var total_amount float64
	for _, k := range cart.Items {
		// fmt.Println(k)
		total_num += k.Num
		total_amount += k.Amount
	}
	sqlstr := "update bookstore.carts set num=?, amount=? where userid=?"
	// fmt.Println("data:", total_num, total_amount, userid)
	_, err := utils.DB.Exec(sqlstr, &total_num, &total_amount, &userid)
	if err != nil {
		fmt.Println("UpdateCart utils.DB.Exec error:", err)
		return
	}
}

func CleanCart(userid int) {
	cart := GetCart(userid)
	sqlstr := "delete from bookstore.cart_items where cart_id = ?"
	utils.DB.Exec(sqlstr, cart.ID)
	UpdateCart(userid)
}