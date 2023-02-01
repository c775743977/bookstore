package dao

import (
	"bookstore1.1/model"
	"bookstore1.1/utils"
	"fmt"
	_"strconv"
)

func AddItem(bookid int, cartid int) {
	if CheckItem(bookid, cartid) {
		ItemAddNum(bookid, cartid)
		return
	}
	fmt.Println("bookid:", bookid)
	fmt.Println("cartid:", cartid)
	sqlstr := "select price from bookstore.books where id = ?"
	row := utils.DB.QueryRow(sqlstr, bookid)
	var price float64
	err := row.Scan(&price)
	if err != nil {
		fmt.Println("AddItem row.Scan error:", err)
		return
	}
	sqlstr = "insert into bookstore.cart_items value(?,?,?,?)"
	_, err = utils.DB.Exec(sqlstr, cartid, bookid, 1, price)
	if err != nil {
		fmt.Println("AddItem utils.DB.Exec error:", err)
		return
	}
}

func CheckItem(bookid int, cartid int) bool {
	// sqlstr := "select book_id from bookstore.cart_items where id = ?"
	// rows, err := utils.DB.Query(sqlstr, cartid)
	// if err != nil {
	// 	fmt.Println("CheckItem utils.DB.Query error:", err)
	// 	return false
	// }
	// var res int
	// for rows.Next() {
	// 	err = rows.Scan(&res)
	// 	if err != nil {
	// 		fmt.Println("CheckItem row.Scan error:", err)
	// 		return false
	// 	}
	// 	if res == bookid {
	// 		return true
	// 	} else {
	// 		return false
	// 	}
	// }
	// return false
	sqlstr := "select num from bookstore.cart_items where cart_id = ? and book_id = ?"
	row := utils.DB.QueryRow(sqlstr, cartid, bookid)
	var res int
	err := row.Scan(&res)
	if err != nil {
		// fmt.Println("CheckItem row.Scan error:", err)
		return false
	}
	return true
}

func ItemAddNum(bookid int, cartid int) {
	sqlstr := "update bookstore.cart_items set amount=amount+amount/num, num=num+1 where book_id = ? and cart_id = ?"
	_, err := utils.DB.Exec(sqlstr, bookid, cartid)
	if err != nil {
		fmt.Println("ItemAddNum utils.DB.Exec error:", err)
		return
	}
}

func ModifyNum(bookid int, cartid int, num int) {
	sqlstr := "update bookstore.cart_items set amount=?*(amount/num), num=? where book_id = ? and cart_id = ?"
	_, err := utils.DB.Exec(sqlstr, num, num, bookid, cartid)
	if err != nil {
		fmt.Println("ItemAddNum utils.DB.Exec error:", err)
		return
	}
}

func GetItems(cartid int) []*model.CartItems {
	// fmt.Println("cartid:", cartid)
	var ci []*model.CartItems
	sqlstr := "select * from bookstore.cart_items where cart_id = ?"
	rows, err := utils.DB.Query(sqlstr, cartid)
	if err != nil {
		fmt.Println("GetItems utils.DB.Query error:", err)
		return nil
	}
	for rows.Next() {
		item := &model.CartItems{}
		book := &model.Book{}
		rows.Scan(&item.CartID, &book.ID, &item.Num, &item.Amount)
		item.Book = GetBook(int(book.ID))
		// fmt.Println("GetItems item:", item)
		ci = append(ci, item)
	}
	return ci
}

func DelItem(bookid int, cartid int) {
	sqlstr := "delete from bookstore.cart_items where book_id = ? and cart_id = ?"
	_, err := utils.DB.Exec(sqlstr, bookid, cartid)
	if err != nil {
		fmt.Println("DelItem utils.DB.Exec error:", err)
		return
	}
}
