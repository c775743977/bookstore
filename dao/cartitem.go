package dao

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"bookstore1.4/model"
	"bookstore1.4/utils"
	"fmt"
	_"strconv"
)

func AddItem(bookid string, cartid string) {
	if CheckItem(bookid, cartid) {
		ItemAddNum(bookid, cartid)
		return
	}
	var data bson.D
	bid, _ := primitive.ObjectIDFromHex(bookid)
	cid, _ := primitive.ObjectIDFromHex(cartid)
	res := utils.C_books.FindOne(utils.Ctx, bson.D{
		{"_id", bid},
	})
	err := res.Decode(&data)
	if err != nil {
		fmt.Println("res.Decode(&data) error:", err)
		return 
	}
	// utils.DBrr.RoundRobin().Model(&model.Book{}).Where("id = ?", bookid).Select("price").Find(&price)
	var ci = bson.D{
		{"cart_id", cid},
		{"book_id", bid},
		{"num", 1},
		{"amount", data[3].Value.(float64)},
	}
	fmt.Println("ci:", ci)
	_, err = utils.C_cartitems.InsertOne(utils.Ctx, ci)
	if err != nil {
		fmt.Println("C_cartitems.InsertOne error:", err)
		return
	}
}

func CheckItem(bookid string, cartid string) bool {
	var item bson.D
	bid, _ := primitive.ObjectIDFromHex(bookid)
	cid, _ := primitive.ObjectIDFromHex(cartid)
	data := utils.C_cartitems.FindOne(utils.Ctx, bson.D{
		{"cart_id", cid},
		{"book_id", bid},
	})
	err := data.Decode(&item)
	if err != nil {
		fmt.Println("data.Decode(&item) error:", err)
		return false
	}
	// utils.DBrr.RoundRobin().Model(&model.CartItem{}).Where("cart_id = ? AND book_id = ?", cartid, bookid).Select("num").Find(&res)
	if item[3].Value.(int) > 0 {
		return true
	} else {
		return false
	}
}

func ItemAddNum(bookid string, cartid string) {
	var data bson.D
	bid, _ := primitive.ObjectIDFromHex(bookid)
	cid, _ := primitive.ObjectIDFromHex(cartid)
	res := utils.C_cartitems.FindOne(utils.Ctx, bson.D{
		{"cart_id", cid},
		{"book_id", bid},
	})
	err := res.Decode(&data)
	if err != nil {
		fmt.Println("res.Decode(&data) error:", err)
		return
	}
	amount := data[4].Value.(float64) + data[4].Value.(float64)/float64(data[3].Value.(int))
	num := data[3].Value.(int) + 1
	_, err = utils.C_cartitems.UpdateOne(utils.Ctx, bson.D{}, bson.D{
		{"$set", bson.D{{"num", num},}},
		{"$set", bson.D{{"amount", amount},}},
	})
	if err != nil {
		fmt.Println("C_cartitems.UpdateOne error:", err)
		return
	}
	// var ci model.CartItem
	// utils.DBrr.RoundRobin().Where("book_id = ? AND  cart_id = ?", bookid, cartid).Find(&ci)
	// ci.Amount = ci.Amount + ci.Amount/float64(ci.Num)
	// ci.Num = ci.Num + 1
	// utils.WDB.Where("book_id = ? AND  cart_id = ?", bookid, cartid).Updates(&ci)
}

func ModifyNum(bookid string, cartid string, num int) {
	var item bson.D
	bid, _ := primitive.ObjectIDFromHex(bookid)
	cid, _ := primitive.ObjectIDFromHex(cartid)
	data := utils.C_cartitems.FindOne(utils.Ctx, bson.D{
		{"cart_id", cid},
		{"book_id", bid},
	})
	err := data.Decode(&item)
	if err != nil {
		fmt.Println("data.Decode(&item) error:", err)
		return 
	}
	amount := (item[4].Value.(float64)/float64(item[3].Value.(int32))) * float64(num)
	_ ,err = utils.C_cartitems.UpdateOne(utils.Ctx, bson.D{{"cart_id", cid},{"book_id", bid},}, bson.D{
		{"$set", bson.D{{"num", item[3].Value.(int32)},}},
		{"$set", bson.D{{"amount", amount},}},
	})
	if err != nil {
		fmt.Println("C_cartitems.UpdateOne error:", err)
		return
	}
}
// func ModifyNum(bookid int, cartid int, num int) {
// 	var ci model.CartItem
// 	// utils.DBrr.RoundRobin().Where("book_id = ? AND  cart_id = ?", bookid, cartid).Find(&ci)
// 	utils.RDB1.Where("book_id = ? AND  cart_id = ?", bookid, cartid).Find(&ci)
// 	fmt.Println("amount=", ci.Amount)
// 	ci.Amount = (ci.Amount / float64(ci.Num)) * float64(num)
// 	ci.Num = num
// 	utils.WDB.Where("book_id = ? AND  cart_id = ?", bookid, cartid).Updates(&ci)
// }

func GetItems(cartid string) []*model.CartItem {
	var cis []*model.CartItem
	var data bson.D
	// utils.WDB.Find(&cis)
	// for _, k := range cis {
	// 	book := GetBook(k.BookID)
	// 	k.Book = book
	// }
	cid, _ := primitive.ObjectIDFromHex(cartid)
	cursor, err := utils.C_cartitems.Find(utils.Ctx, bson.D{{"cart_id", cid}})
	if err != nil {
		fmt.Println("C_cartitems.Find error:", err)
		return nil
	}
	for cursor.Next(utils.Ctx) {
		err = cursor.Decode(&data)
		if err != nil {
			fmt.Println("cursor.Decode(&data) error:", err)
			return nil
		}
		ci := model.CartItem{
			CartID : data[1].Value.(primitive.ObjectID).Hex(),
			BookID : data[2].Value.(primitive.ObjectID).Hex(),
			Num : data[3].Value.(int32),
			Amount : data[4].Value.(float64),
		}
		book := GetBook(ci.BookID)
		ci.Book = book
		cis = append(cis, &ci)
	}
	return cis
}

func DelItem(bookid string, cartid string) {
	// utils.WDB.Where("book_id = ? AND cart_id = ?", bookid, cartid).Delete(&model.CartItem{})
	bid, _ := primitive.ObjectIDFromHex(bookid)
	cid, _ := primitive.ObjectIDFromHex(cartid)
	_, err := utils.C_cartitems.DeleteOne(utils.Ctx, bson.D{
		{"book_id", bid},
		{"cart_id", cid},
	})
	if err != nil {
		fmt.Println("C_cartitems.DeleteOne error:", err)
		return
	}
}