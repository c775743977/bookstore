package main

import (
	"net/http"
	_"html/template"
	"fmt"
	"Hello_golang/bookstore/controller"
)


func main() {
	http.HandleFunc("/", controller.Handler) //主页
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static")))) //设置静态资源
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages")))) //设置静态资源
	http.HandleFunc("/login", controller.LoginHandler) //登录页面
	http.HandleFunc("/regist", controller.ResgisterHandler) //注册页面
	// http.HandleFunc("/book_manager", controller.GetBooksHandler) //图书管理页面
	// http.HandleFunc("/add_book", controller.AddBookHandler) //添加图书页面
	http.HandleFunc("/deletebook", controller.DelBookHandler) //删除图书页面
	http.HandleFunc("/tobook_alter_add", controller.ToAddOrAlterBook) //前往添加或更改图书页面
	http.HandleFunc("/book_alter_add", controller.AddOrAlterBook) //提交添加或更改操作
	http.HandleFunc("/book_manager", controller.GetPageHandler) //带有页数的图书管理页面
	// http.HandleFunc("/GetPageByPrice", controller.GetPageByPriceHandler)
	http.HandleFunc("/logout", controller.LogOut)  //注销
	http.HandleFunc("/cart", controller.ShowCart) //显示购物车
	http.HandleFunc("/cart2", controller.ShowCart2)
	http.HandleFunc("/CartAlter", controller.CartAlter)
	http.HandleFunc("/AddBookToCart", controller.AddBookToCart) //添加图书到购物车
	http.HandleFunc("/delete_item", controller.DelItemHandler) //在购物车中删除购物项
	http.HandleFunc("/CleanUp", controller.CleanCartHandler) //清空购物车
	http.HandleFunc("/checkout", controller.SubmitOrder) //提交订单
	http.HandleFunc("/MyOrders", controller.ShowOrders)  //我的订单页面
	http.HandleFunc("/showOrderItems", controller.ShowOrderItems) //订单详情页面
	http.HandleFunc("/cancelOrder", controller.DelOrderHandler)  //删除订单
	http.HandleFunc("/pay", controller.PayHandler) //付款
	http.HandleFunc("/signed", controller.SignHandler) //付款
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println("http.ListenAndServe error:", err)
	}
}