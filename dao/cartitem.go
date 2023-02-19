package dao

import (
	"bookstore1.4/model"
	"bookstore1.4/utils"
	_"fmt"
	_"strconv"
)

func AddItem(bookid int, cartid int) {
	if CheckItem(bookid, cartid) {
		ItemAddNum(bookid, cartid)
		return
	}
	var price float64
	utils.DBrr.RoundRobin().Model(&model.Book{}).Where("id = ?", bookid).Select("price").Find(&price)
	var ci = model.CartItem{
		CartID : cartid,
		BookID : bookid,
		Num : 1,
		Amount : price,
	}
	utils.WDB.Create(&ci)
}

func CheckItem(bookid int, cartid int) bool {
	var res int
	utils.DBrr.RoundRobin().Model(&model.CartItem{}).Where("cart_id = ? AND book_id = ?", cartid, bookid).Select("num").Find(&res)
	if res > 0 {
		return true
	} else {
		return false
	}
}

func ItemAddNum(bookid int, cartid int) {
	var ci model.CartItem
	utils.DBrr.RoundRobin().Where("book_id = ? AND  cart_id = ?", bookid, cartid).Find(&ci)
	ci.Amount = ci.Amount + ci.Amount/float64(ci.Num)
	ci.Num = ci.Num + 1
	utils.WDB.Where("book_id = ? AND  cart_id = ?", bookid, cartid).Updates(&ci)
}

func ModifyNum(bookid int, cartid int, num int) {
	var ci model.CartItem
	utils.DBrr.RoundRobin().Where("book_id = ? AND  cart_id = ?", bookid, cartid).Find(&ci)
	ci.Amount = (ci.Amount / float64(ci.Num)) * float64(num)
	ci.Num = num
	utils.WDB.Where("book_id = ? AND  cart_id = ?", bookid, cartid).Updates(&ci)
}

func GetItems(cartid int) []*model.CartItem {
	var cis []*model.CartItem
	utils.WDB.Find(&cis)
	for _, k := range cis {
		book := GetBook(k.BookID)
		k.Book = book
	}
	return cis
}

func DelItem(bookid int, cartid int) {
	utils.WDB.Where("book_id = ? AND cart_id = ?", bookid, cartid).Delete(&model.CartItem{})
}