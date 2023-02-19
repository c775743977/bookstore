package dao

import (
	"bookstore1.4/model"
	"bookstore1.4/utils"
	"fmt"
	_"strconv"
	"time"
)

func GetTime() string {
	t := time.Now()
	y := fmt.Sprint(t.Year())
	m := fmt.Sprint(int(t.Month()))
	d := fmt.Sprint(t.Day())
	h := fmt.Sprint(t.Hour())
	mi := fmt.Sprint(t.Minute())
	s := fmt.Sprint(t.Second())
	time := y+"-"+m+"-"+d+" "+h+":"+mi+":"+s
	return time
}

func CreateOrder(cart *model.Cart) *model.Order {
	var order model.Order
	order.ID = utils.CreateOrderID()
	order.TotalCount = int64(cart.Num)
	order.TotalAmount = cart.Amount
	order.Status = 0
	order.UserID = int64(cart.UserID)
	order.CreateTime = GetTime()
	utils.WDB.Create(&order)
	return &order
}

func AddOrderItems(cart *model.Cart, orderid string) []*model.OrderItem {
	var items []*model.OrderItem
	for _, k := range cart.Items {
		var item model.OrderItem
		item.OrderID = orderid
		item.Num = int64(k.Num)
		item.Amount = k.Amount
		item.Title = k.Book.Title
		item.Author = k.Book.Author
		item.Price = k.Book.Price
		items = append(items, &item)
		utils.WDB.Create(&item)
	}
	return items
}

func GetOrders(userid int) []*model.Order {
	var orders []*model.Order
	utils.DBrr.RoundRobin().Where("user_id = ?", userid).Find(&orders)
	return orders
}

func GetOrderItems(orderid string) []*model.OrderItem {
	var items []*model.OrderItem
	utils.DBrr.RoundRobin().Where("order_id = ?", orderid).Find(&items)
	return items
}

func DelOrder(orderid string) {
	utils.WDB.Where("order_id = ?", orderid).Delete(&model.OrderItem{})
	utils.WDB.Where("id = ?", orderid).Delete(&model.Order{})
}

func Pay(orderid string) {
	utils.WDB.Where("id = ?", orderid).Select("status").Updates(model.Order{Status : 1})
}

func Sign(orderid string) {
	utils.WDB.Where("id = ?", orderid).Select("status").Updates(model.Order{Status : 4})
}

func Deliver(orderid string) {
	utils.WDB.Where("id = ?", orderid).Select("status").Updates(model.Order{Status : 3})
}

func TakeOrder(orderid string) {
	utils.WDB.Where("id = ?", orderid).Select("status").Updates(model.Order{Status : 2})
}

func GetAllOrders() []*model.Order {
	var orders []*model.Order
	utils.DBrr.RoundRobin().Find(&orders)
	for _, k := range orders {
		user := GetUserByID(int(k.UserID))
		k.UserName = user.Name
	}
	return orders
}