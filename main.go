package main

import (
	"github.com/gin-gonic/gin"
	"bookstore1.4/controller"
	"bookstore1.4/utils"
	"html/template"
)

func ErrInfo() (info string) {
	info = "失败！不允许提交空白数据"
	return info
}

func main() {
	defer utils.Cancel()
	r := gin.Default()
	//定义一个输出错误信息的函数，所有加载的页面都能使用
	r.SetFuncMap(template.FuncMap{
		"ErrInfo" : ErrInfo,
	})
	r.LoadHTMLGlob("pages/**/*")
	r.Static("/static", "./static")
	r.Static("/pages", "./pages")
	r.GET("/home", controller.Handler) //主页
	r.POST("/home", controller.Handler) //主页
	r.GET("/login", func(c *gin.Context) { //访问登录页面
		c.HTML(200, "user/login.html", "")
	})
	r.POST("/login", controller.LoginHandler) //处理登录请求
	r.GET("/regist", func(c *gin.Context) { //访问注册页面
		c.HTML(200, "user/regist.html", "")
	})
	r.POST("/regist", controller.RegistHandler) //处理注册请求
	r.GET("/manager", func(c *gin.Context) { //访问管理员页面
		c.HTML(200, "manager/manager.html", nil)
	})
	r.GET("/book_manager", controller.GetPageHandler) //图书管理页面
	r.GET("/deletebook", controller.DelBookHandler) //执行删除图书
	r.GET("/book_alter_add", controller.ToAddOrAlterBook) //访问修改或添加图书页面，仅访问页面还没进行操作
	r.POST("/book_alter_add", controller.AddOrAlterBook) //执行修改或添加图书操作
	r.GET("/user_manager", controller.UsersManageHandler) //用户管理页面
	r.GET("/user_delete", controller.DelUserHandler) //执行删除用户
	r.GET("/user_alter", controller.ToAlterUserHandler) //访问修改或添加用户页面，仅访问页面还没进行操作
	r.POST("/user_alter", controller.AlterUserHandler) //执行修改或添加用户操作
	r.GET("/logout", controller.LogoutHandler) //注销
	r.GET("/cart", controller.ShowCart)  //购物车页面
	r.GET("/delete_item", controller.DelItemHandler)  //删除购物项
	r.GET("/CleanUp", controller.CleanCartHandler) //清空购物车
	// r.GET("/cart2", controller.ShowCart2) 
	r.POST("/AddBookToCart", controller.AddBookToCart) //从主页添加图书至购物车
	r.POST("/CartAlter", controller.CartAlter)
	r.GET("/checkout", controller.SubmitOrder) //提交订单
	r.GET("/MyOrders", controller.ShowOrders)  //我的订单页面
	r.GET("/showOrderItems", controller.ShowOrderItems) //订单详情页面
	r.GET("/cancelOrder", controller.DelOrderHandler)  //删除订单
	r.GET("/pay", controller.PayHandler) //付款
	r.GET("/signed", controller.SignHandler) //签收
	r.GET("/order_manager", controller.GetAllOrdersHandler) //订单管理界面
	r.GET("/deliver", controller.DeliverHandler) //发货
	r.GET("/McancelOrder", controller.MDelOrderHandler) //管理员删除订单
	r.GET("/takeOrder", controller.TakeOrderHandler) //接收订单
	r.Run()
}