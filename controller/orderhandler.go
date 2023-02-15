package controller

import (
	"bookstore1.4/dao"
	"bookstore1.4/model"
	"github.com/gin-gonic/gin"
	_"fmt"
	"net/http"
	_"html/template"
	_"strconv"
)

func SubmitOrder(c *gin.Context) {
	cookie, err := c.Cookie("user")
	if err != nil {
		c.String(400, "请先登录...")
		return
	}
	sess := dao.GetSession(cookie)
	cart := dao.GetCart(sess.UserID)
	if cart.Num == 0 {
		return
	}
	order := dao.CreateOrder(cart)
	dao.AddOrderItems(cart, order.ID)
	dao.CleanCart(sess.UserID)
	c.HTML(http.StatusOK, "cart/checkout.html", order)
}

func ShowOrders(c *gin.Context) {
	cookie, err := c.Cookie("user")
	if err != nil {
		c.String(400, "请先登录...")
		return
	}
	sess := dao.GetSession(cookie)
	orders := dao.GetOrders(sess.UserID)
	user := &model.User{
		ID : sess.UserID,
		Name : sess.Username,
		Orders : orders,
	}
	c.HTML(http.StatusOK, "order/order.html", user)
}


func ShowOrderItems(c *gin.Context) {
	orderid := c.Query("orderID")
	items := dao.GetOrderItems(orderid)
	c.HTML(http.StatusOK, "order/order_info.html", items)
}

func DelOrderHandler(c *gin.Context) {
	orderid := c.Query("orderID")
	dao.DelOrder(orderid)
	ShowOrders(c)
}

func PayHandler(c *gin.Context) {
	orderid := c.Query("orderID")
	dao.Pay(orderid)
	ShowOrders(c)
}

func SignHandler(c *gin.Context) {
	orderid := c.Query("orderID")
	dao.Sign(orderid)
	ShowOrders(c)
}

func DeliverHandler(c *gin.Context) {
	orderid := c.Query("orderID")
	dao.Deliver(orderid)
	GetAllOrdersHandler(c)
}

func TakeOrderHandler(c *gin.Context) {
	orderid := c.Query("orderID")
	dao.TakeOrder(orderid)
	GetAllOrdersHandler(c)
}

func GetAllOrdersHandler(c *gin.Context) {
	orders := dao.GetAllOrders()
	c.HTML(http.StatusOK, "manager/order_manager.html", orders)
}

func MDelOrderHandler(c *gin.Context) {
	orderid := c.Query("orderID")
	dao.DelOrder(orderid)
	GetAllOrdersHandler(c)
}