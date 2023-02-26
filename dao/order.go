package dao

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	order.TotalCount = int32(cart.Num)
	order.TotalAmount = cart.Amount
	order.Status = 0
	order.UserID = cart.UserID
	order.CreateTime = GetTime()
	userid, _ := primitive.ObjectIDFromHex(order.UserID)
	var data = bson.D{
		{"_id", order.ID},
		{"create_time", order.CreateTime},
		{"total_count", order.TotalCount},
		{"total_amount", order.TotalAmount},
		{"status", order.Status},
		{"user_id", userid},
	}
	_, err := utils.C_orders.InsertOne(utils.Ctx, data)
	if err != nil {
		fmt.Println("C_orders.InsertOne error:", err)
		return nil
	}
	return &order
}

func AddOrderItems(cart *model.Cart, orderid string) []*model.OrderItem {
	var items []*model.OrderItem
	for _, k := range cart.Items {
		var item model.OrderItem
		item.OrderID = orderid
		item.Num = int32(k.Num)
		item.Amount = k.Amount
		item.Title = k.Book.Title
		item.Author = k.Book.Author
		item.Price = k.Book.Price
		items = append(items, &item)
		// utils.WDB.Create(&item)
		var data = bson.D {
			{"order_id", item.OrderID},
			{"num", item.Num},
			{"amount", item.Amount},
			{"title", item.Title},
			{"author", item.Author},
			{"price", item.Price},
		}
		_, err := utils.C_orderitems.InsertOne(utils.Ctx, data)
		if err != nil {
			fmt.Println("C_orderitems.InsertOne error:", err)
			return nil
		}
	}
	return items
}

func GetOrders(userid string) []*model.Order {
	var orders []*model.Order
	uid, _ := primitive.ObjectIDFromHex(userid)
	// utils.DBrr.RoundRobin().Where("user_id = ?", userid).Find(&orders)
	cursor, err := utils.C_orders.Find(utils.Ctx, bson.D{{"user_id", uid}})
	if err != nil {
		fmt.Println("C_orders.Find error:", err)
		return nil
	}
	for cursor.Next(utils.Ctx) {
		var data bson.D
		err = cursor.Decode(&data)
		if err != nil {
			fmt.Println("cursor.Decode(&data) error:", err)
			return nil
		}
		var order = model.Order{
			ID : data[0].Value.(string),
			CreateTime : data[1].Value.(string),
			TotalCount : data[2].Value.(int32),
			TotalAmount : data[3].Value.(float64),
			Status : data[4].Value.(int32),
			UserID : data[5].Value.(primitive.ObjectID).Hex(),
		}
		orders = append(orders, &order)
	}
	return orders
}

func GetOrderItems(orderid string) []*model.OrderItem {
	var items []*model.OrderItem
	var data bson.D
	// utils.DBrr.RoundRobin().Where("order_id = ?", orderid).Find(&items)
	cursor, err := utils.C_orderitems.Find(utils.Ctx, bson.D{{"order_id", orderid},})
	if err != nil {
		fmt.Println("C_orderitems.Find error:", err)
		return nil
	}
	for cursor.Next(utils.Ctx) {
		err = cursor.Decode(&data)
		if err != nil {
			fmt.Println("cursor.Decode(&data) error:", err)
			return nil
		}
		var item = model.OrderItem{
			OrderID : data[1].Value.(string),
			Num : data[2].Value.(int32),
			Amount : data[3].Value.(float64),
			Title : data[4].Value.(string),
			Author : data[5].Value.(string),
			Price : data[6].Value.(float64),
		}
		items = append(items, &item)
	}
	return items
}

func DelOrder(orderid string) {
	// utils.WDB.Where("order_id = ?", orderid).Delete(&model.OrderItem{})
	// utils.WDB.Where("id = ?", orderid).Delete(&model.Order{})
	_, err := utils.C_orderitems.DeleteMany(utils.Ctx, bson.D{{"order_id", orderid},})
	if err != nil {
		fmt.Println("C_orderitems.DeleteMany error:", err)
		return
	}
	_, err = utils.C_orders.DeleteMany(utils.Ctx, bson.D{{"_id", orderid},})
	if err != nil {
		fmt.Println("C_orders.DeleteMany error:", err)
		return
	}
}

func Pay(orderid string) {
	// utils.WDB.Where("id = ?", orderid).Select("status").Updates(model.Order{Status : 1})
	_, err := utils.C_orders.UpdateOne(utils.Ctx, bson.D{{"_id", orderid},}, bson.D{{"$set", bson.D{{"status", 1},}}})
	if err != nil {
		fmt.Println("C_orders.UpdateOne error:", err)
		return
	}
}

func Sign(orderid string) {
	// utils.WDB.Where("id = ?", orderid).Select("status").Updates(model.Order{Status : 4})
	_, err := utils.C_orders.UpdateOne(utils.Ctx, bson.D{{"_id", orderid},}, bson.D{{"$set", bson.D{{"status", 4},}}})
	if err != nil {
		fmt.Println("C_orders.UpdateOne error:", err)
		return
	}
}

func Deliver(orderid string) {
	// utils.WDB.Where("id = ?", orderid).Select("status").Updates(model.Order{Status : 3})
	_, err := utils.C_orders.UpdateOne(utils.Ctx, bson.D{{"_id", orderid},}, bson.D{{"$set", bson.D{{"status", 3},}}})
	if err != nil {
		fmt.Println("C_orders.UpdateOne error:", err)
		return
	}
}

func TakeOrder(orderid string) {
	// utils.WDB.Where("id = ?", orderid).Select("status").Updates(model.Order{Status : 2})
	_, err := utils.C_orders.UpdateOne(utils.Ctx, bson.D{{"_id", orderid},}, bson.D{{"$set", bson.D{{"status", 2},}}})
	if err != nil {
		fmt.Println("C_orders.UpdateOne error:", err)
		return
	}
}

func GetAllOrders() []*model.Order {
	var orders []*model.Order
	// utils.DBrr.RoundRobin().Find(&orders)
	// for _, k := range orders {
	// 	user := GetUserByID(int(k.UserID))
	// 	k.UserName = user.Name
	// }
	cursor, err := utils.C_orders.Find(utils.Ctx, bson.D{})
	if err != nil {
		fmt.Println("C_orders.Find error:", err)
		return nil
	}
	var data bson.D
	for cursor.Next(utils.Ctx) {
		err = cursor.Decode(&data)
		if err != nil {
			fmt.Println("cursor.Decode(&data) error:", err)
			return nil
		}
		var order = model.Order{
			ID : data[0].Value.(string),
			CreateTime : data[1].Value.(string),
			TotalCount : data[2].Value.(int32),
			TotalAmount : data[3].Value.(float64),
			Status : data[4].Value.(int32),
			UserID : data[5].Value.(primitive.ObjectID).Hex(),
		}
		orders = append(orders, &order)
	}
	return orders
}