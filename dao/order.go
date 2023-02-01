package dao

import (
	"bookstore1.2/model"
	"bookstore1.2/utils"
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
	order := &model.Order{}
	order.OrderID = utils.CreateOrderID()
	order.Num = int64(cart.Num)
	order.Amount = cart.Amount
	order.Status = 0
	order.UserID = int64(cart.UserID)
	order.CreateTime = GetTime()
	sqlstr := "insert into bookstore.orders value(?,?,?,?,?,?)"
	_, err := utils.DB.Exec(sqlstr, order.OrderID, order.CreateTime, order.Num, order.Amount, order.Status, order.UserID)
	if err != nil {
		fmt.Println("CreateOrder utils.DB.Exec error:", err)
		return nil
	}
	return order
}

func AddOrderItems(cart *model.Cart, orderid string) []*model.OrderItems {
	var items []*model.OrderItems
	sqlstr := "insert into bookstore.orderitems value(?,?,?,?,?,?)"
	for _, k := range cart.Items {
		item := &model.OrderItems{}
		item.OrderID = orderid
		item.Num = int64(k.Num)
		item.Amount = k.Amount
		item.Title = k.Book.Title
		item.Author = k.Book.Author
		item.Price = k.Book.Price
		items = append(items, item)
		_, err := utils.DB.Exec(sqlstr, item.OrderID, item.Num, item.Amount, item.Title, item.Author, item.Price)
		if err != nil {
			fmt.Println("AddOrderItems utils.DB.Exec error:", err)
			return nil
		}
	}
	return items
}

func GetOrders(userid int) []*model.Order {
	var orders []*model.Order
	sqlstr := "select * from bookstore.orders where user_id = ?"
	rows, err := utils.DB.Query(sqlstr, userid)
	if err != nil {
		fmt.Println("GetOrder utils.DB.Query error:", err)
		return nil
	}
	for rows.Next() {
		order := &model.Order{}
		rows.Scan(&order.OrderID, &order.CreateTime, &order.Num, &order.Amount, &order.Status, &order.UserID)
		orders = append(orders, order)
	}
	return orders
}

func GetOrderItems(orderid string) []*model.OrderItems {
	var items []*model.OrderItems
	sqlstr := "select * from bookstore.orderitems where order_id = ?"
	rows, err := utils.DB.Query(sqlstr, orderid)
	if err != nil {
		fmt.Println("GetOrder utils.DB.Query error:", err)
		return nil
	}
	for rows.Next() {
		item := &model.OrderItems{}
		rows.Scan(&item.OrderID, &item.Num, &item.Amount, &item.Title, &item.Author, &item.Price)
		items = append(items, item)
	}
	return items
}

func DelOrder(orderid string) {
	sqlstr := "delete from bookstore.orderitems where order_id = ?"
	_, err := utils.DB.Exec(sqlstr, orderid)
	if err != nil {
		fmt.Println("DelOrder utils.DB.Exec1 error:", err)
		return
	}
	sqlstr = "delete from bookstore.orders where id = ?"
	_, err = utils.DB.Exec(sqlstr, orderid)
	if err != nil {
		fmt.Println("DelOrder utils.DB.Exec2 error:", err)
		return
	}
}

func Pay(orderid string) {
	sqlstr := "update bookstore.orders set status=1 where id=?"
	_, err := utils.DB.Exec(sqlstr, orderid)
	if err != nil {
		fmt.Println("Pay utils.DB.Exec error:", err)
		return
	}
}

func Sign(orderid string) {
	sqlstr := "update bookstore.orders set status=4 where id=?"
	_, err := utils.DB.Exec(sqlstr, orderid)
	if err != nil {
		fmt.Println("Sign utils.DB.Exec error:", err)
		return
	}
}

func Deliver(orderid string) {
	sqlstr := "update bookstore.orders set status=3 where id=?"
	_, err := utils.DB.Exec(sqlstr, orderid)
	if err != nil {
		fmt.Println("Sign utils.DB.Exec error:", err)
		return
	}
}

func TakeOrder(orderid string) {
	sqlstr := "update bookstore.orders set status=2 where id=?"
	_, err := utils.DB.Exec(sqlstr, orderid)
	if err != nil {
		fmt.Println("Sign utils.DB.Exec error:", err)
		return
	}
}

func GetAllOrders() []*model.Order {
	var orders []*model.Order
	sqlstr := "select * from bookstore.orders"
	rows, err := utils.DB.Query(sqlstr)
	if err != nil {
		fmt.Println("GetOrder utils.DB.Query error:", err)
		return nil
	}
	for rows.Next() {
		order := &model.Order{}
		rows.Scan(&order.OrderID, &order.CreateTime, &order.Num, &order.Amount, &order.Status, &order.UserID)
		order.UserName = GetUserByID(int(order.UserID))
		orders = append(orders, order)
	}
	return orders
}