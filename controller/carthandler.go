package controller

import (
	"Hello_golang/bookstore/dao"
	"fmt"
	"net/http"
	"html/template"
	"strconv"
)

func ShowCart(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("user")
	username := ""
	userid := 0
	if cookie != nil {
		username, userid = dao.GetSession(cookie.Value)
	} else {
		return
	}
	cart := dao.GetCart(userid)
	cart.UserName = username
	t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
	t.Execute(w, cart)
}

func ShowCart2(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("user")
	itemID, _ := strconv.ParseInt(r.FormValue("item"), 0, 0)
	username := ""
	userid := 0
	if cookie != nil {
		username, userid = dao.GetSession(cookie.Value)
	}
	cart := dao.GetCart(userid)
	cart.UserName = username
	for _, k := range cart.Items {
		// fmt.Println("data:", k)
		if itemID == k.Book.ID {
			k.IsThis = true
			break
		}
	}
	t := template.Must(template.ParseFiles("views/pages/cart/cart2.html"))
	t.Execute(w, cart)
}

func AddBookToCart(w http.ResponseWriter, r *http.Request) {  //从主页添加图书到购物车
	bookid, _ := strconv.ParseInt(r.FormValue("bookId"), 0, 0)
	fmt.Println("BOOKID:", bookid)
	cookie, _ := r.Cookie("user")
	userid := 0
	if cookie != nil {
		_, userid = dao.GetSession(cookie.Value)
	}
	if userid == 0 {
		return
	}
	cart := dao.GetCart(userid)
	if cart == nil {
		Handler(w, r)
		return
	}
	dao.AddItem(int(bookid), cart.ID)
	dao.UpdateCart(userid)
	Handler(w, r)
}

func DelItemHandler(w http.ResponseWriter, r *http.Request) {  //删除购物车某项商品
	cookie, _ := r.Cookie("user")
	username := ""
	userid := 0
	if cookie != nil {
		username, userid = dao.GetSession(cookie.Value)
	} else {
		return
	}
	cart := dao.GetCart(userid)
	bookid, _ := strconv.ParseInt(r.FormValue("bookID"), 0, 0)
	dao.DelItem(int(bookid), cart.ID)
	dao.UpdateCart(userid)
	cart = dao.GetCart(userid)
	// for i, k := range cart.Items {
	// 	if k.Book.ID == bookid {
	// 		cart.Items = append(cart.Items[:i], cart.Items[i+1:]...)
	// 	}
	// }
	cart.UserName = username
	t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
	t.Execute(w, cart)
}

func CleanCartHandler(w http.ResponseWriter, r *http.Request) { //清空购物车
	cookie, _ := r.Cookie("user")
	username := ""
	userid := 0
	if cookie != nil {
		username, userid = dao.GetSession(cookie.Value)
	} else {
		return
	}
	dao.CleanCart(userid)
	cart := dao.GetCart(userid)
	cart.UserName = username
	t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
	t.Execute(w, cart)
}

func CartAlter(w http.ResponseWriter, r *http.Request) { //修改购物车商品数量
	cookie, _ := r.Cookie("user")
	booknum, _ := strconv.ParseInt(r.PostFormValue("BookNum"), 0, 0)
	bookid, _ := strconv.ParseInt(r.FormValue("BookID"), 0, 0)
	username := ""
	userid := 0
	if cookie != nil {
		username, userid = dao.GetSession(cookie.Value)
	} else {
		return
	}
	cart := dao.GetCart(userid)
	dao.ModifyNum(int(bookid), cart.ID, int(booknum))
	dao.UpdateCart(userid)
	cart = dao.GetCart(userid)
	cart.UserName = username
	t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
	t.Execute(w, cart)
}