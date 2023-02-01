package controller

import (
	"bookstore1.2/dao"
	"bookstore1.2/model"
	_"fmt"
	"net/http"
	"html/template"
	_"strconv"
)

func SubmitOrder(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("user")
	username := ""
	userid := 0
	if cookie != nil {
		username, userid = dao.GetSession(cookie.Value)
	}
	cart := dao.GetCart(userid)
	if cart.Num == 0 {
		return
	}
	order := dao.CreateOrder(cart)
	order.UserName = username
	dao.AddOrderItems(cart, order.OrderID)
	dao.CleanCart(userid)
	t := template.Must(template.ParseFiles("views/pages/cart/checkout.html"))
	t.Execute(w, order)
}

func ShowOrders(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("user")
	username := ""
	userid := 0
	if cookie != nil {
		username, userid = dao.GetSession(cookie.Value)
	}
	orders := dao.GetOrders(userid)
	user := &model.User{
		ID : userid,
		Name : username,
		Orders : orders,
	}
	t := template.Must(template.ParseFiles("views/pages/order/order.html"))
	t.Execute(w, user)
}

func ShowOrderItems(w http.ResponseWriter, r *http.Request) {
	orderid := r.FormValue("orderID")
	items := dao.GetOrderItems(orderid)
	t := template.Must(template.ParseFiles("views/pages/order/order_info.html"))
	t.Execute(w, items)
}

func DelOrderHandler(w http.ResponseWriter, r *http.Request) {
	orderid := r.FormValue("orderID")
	dao.DelOrder(orderid)
	ShowOrders(w, r)
}

func PayHandler(w http.ResponseWriter, r *http.Request) {
	orderid := r.FormValue("orderID")
	dao.Pay(orderid)
	ShowOrders(w, r)
}

func SignHandler(w http.ResponseWriter, r *http.Request) {
	orderid := r.FormValue("orderID")
	dao.Sign(orderid)
	ShowOrders(w, r)
}

func DeliverHandler(w http.ResponseWriter, r *http.Request) {
	orderid := r.FormValue("orderID")
	dao.Deliver(orderid)
	GetAllOrdersHandler(w, r)
}

func TakeOrderHandler(w http.ResponseWriter, r *http.Request) {
	orderid := r.FormValue("orderID")
	dao.TakeOrder(orderid)
	GetAllOrdersHandler(w, r)
}

func GetAllOrdersHandler(w http.ResponseWriter, r *http.Request) {
	orders := dao.GetAllOrders()
	t := template.Must(template.ParseFiles("views/pages/manager/order_manager.html"))
	t.Execute(w, orders)
}

func MDelOrderHandler(w http.ResponseWriter, r *http.Request) {
	orderid := r.FormValue("orderID")
	dao.DelOrder(orderid)
	GetAllOrdersHandler(w, r)
}