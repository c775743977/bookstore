package dao

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"bookstore1.4/model"
	"bookstore1.4/utils"
	"strconv"
	"fmt"
)

func GetBooks()  []bson.D {
	var books []bson.D
	cursor, err := utils.C_books.Find(utils.Ctx, bson.D{})
	if err != nil {
		fmt.Println("books.Find error:", err)
		return nil
	}
	err = cursor.Decode(&books)
	if err != nil {
		fmt.Println("Decode(&books) error:", err)
		return nil
	}
	return books
}

func GetPage(bpageno string) (page *model.Page) {  //获取一页的图书数据
	pageno, _ := strconv.ParseInt(bpageno, 10, 0) //由pageno是从url的请求数据里读取的，所以默认是string类型
	page = &model.Page{}
	page.TotalBooks, _ = utils.C_books.CountDocuments(utils.Ctx, bson.D{})
	// utils.DBrr.RoundRobin().Model(&model.Book{}).Count(&page.TotalBooks)
	page.PageNo = pageno
	page.PageSize = 4
	if page.TotalBooks % page.PageSize == 0 {
		page.TotalPages = page.TotalBooks / page.PageSize
	} else {
		page.TotalPages = page.TotalBooks / page.PageSize + 1
	}

	var books []*model.Book
	var data bson.D
	opt := options.Find().SetLimit(4).SetSkip((page.PageNo-1)*page.PageSize)
	cursor, _ := utils.C_books.Find(utils.Ctx, bson.D{}, opt)
	for cursor.Next(utils.Ctx) {
		cursor.Decode(&data)
		var book = model.Book {
			ID : data[0].Value.(primitive.ObjectID).Hex(),
			Title : data[1].Value.(string),
			Author : data[2].Value.(string),
			Price : data[3].Value.(float64),
			Sales : data[4].Value.(int32),
			Stock : data[5].Value.(int32),
			Img_path : data[6].Value.(string),
		}
		books = append(books, &book)
	}
	// utils.DBrr.RoundRobin().Limit(int(page.PageSize)).Offset(int((page.PageNo-1)*page.PageSize)).Find(&books)
	page.Books = books
	return page
}

func GetPageByPrice(bpageno string, bmax string, bmin string) (page *model.Page) { //按照价格搜索结果的一页数据
	max, _ := strconv.ParseFloat(bmax, 64)
	min, _ := strconv.ParseFloat(bmin, 64)
	pageno, _ := strconv.ParseInt(bpageno, 10, 0)
	page = &model.Page{}
	page.TotalBooks, _ = utils.C_books.CountDocuments(utils.Ctx, bson.D{
		{"price", bson.D{
			{"$gt", min},
			{"$lt", max},
		}},
	})
	// utils.DBrr.RoundRobin().Model(&model.Book{}).Where("price between ? and ?", min, max).Count(&page.TotalBooks)
	page.PageNo = pageno
	page.PageSize = 4
	if page.TotalBooks % page.PageSize == 0 {
		page.TotalPages = page.TotalBooks / page.PageSize
	} else {
		page.TotalPages = page.TotalBooks / page.PageSize + 1
	}
	var books []*model.Book
	var data bson.D
	opt := options.Find().SetHint("price_1").SetMax(bson.D{{"price", max},}).SetMin(bson.D{{"price", min}}).SetLimit(4).SetSkip((page.PageNo-1)*page.PageSize)
	cursor, _ := utils.C_books.Find(utils.Ctx, bson.D{}, opt)
	for cursor.Next(utils.Ctx) {
		cursor.Decode(&data)
		var book = model.Book {
			ID : data[0].Value.(primitive.ObjectID).Hex(),
			Title : data[1].Value.(string),
			Author : data[2].Value.(string),
			Price : data[3].Value.(float64),
			Sales : data[4].Value.(int32),
			Stock : data[5].Value.(int32),
			Img_path : "static/img/default.jpg",
		}
		books = append(books, &book)
	}
	// utils.DBrr.RoundRobin().Limit(int(page.PageSize)).Offset(int((page.PageNo-1)*page.PageSize)).Where("price between ? and ?", min, max).Find(&books)
	page.Books = books
	page.MaxPrice = max
	page.MinPrice = min
	return page
} 

func AddBook(book *model.Book) { //添加图书
	var data bson.D = bson.D{
		{"title", book.Title},
		{"author", book.Author},
		{"price", book.Price},
		{"sales", book.Sales},
		{"stock", book.Stock},
		{"img_path", "static/img/default.jpg"},
	}
	_, err := utils.C_books.InsertOne(utils.Ctx, data)
	if err != nil {
		fmt.Println("utils.C_books.InsertOne error:", err)
		return
	}
}

func CheckBook(title string, author string) bool { //检查图书是否存在
	var res bson.D
	cursor, err := utils.C_books.Find(utils.Ctx, bson.D{
		{"author", author},
	})
	if err != nil {
		fmt.Println("C_books.Find error:", err)
		return true
	}
	for cursor.Next(utils.Ctx) {
		err = cursor.Decode(&res)
		if err != nil {
			fmt.Println("cursor.Decode(&res) error:", err)
			return true
		}
		if res[1].Value == title {
			return true
		} 
	}
	return false
}

func DelBook(id string) { //删除图书
	bid, _ := primitive.ObjectIDFromHex(id)
	_, err := utils.C_books.DeleteOne(utils.Ctx, bson.D{
		{"_id", bid},
	})
	if err != nil {
		fmt.Println("C_books.DeleteOne error:", err)
		return
	}
}

func GetBook(id string)  *model.Book { //根据输入ID获取对应图书
	var data bson.D
	bookid, _ := primitive.ObjectIDFromHex(id)
	res := utils.C_books.FindOne(utils.Ctx, bson.D{
		{"_id", bookid},
	})
	err := res.Decode(&data)
	if err != nil {
		fmt.Println("res.Decode(&data) error:", err)
		return nil
	}
	var book = model.Book {
		ID : data[0].Value.(primitive.ObjectID).Hex(),
		Title : data[1].Value.(string),
		Author : data[2].Value.(string),
		Price : data[3].Value.(float64),
		Sales : data[4].Value.(int32),
		Stock : data[5].Value.(int32),
		Img_path : data[6].Value.(string),
	}
	return &book
}

func AlterBook(book *model.Book) {  //更新图书信息
	var data = bson.D{
		{"title", book.Title},
		{"author", book.Author},
		{"price", book.Price},
		{"sales", book.Sales},
		{"stock", book.Stock},
		{"img_path", "static/img/default.jpg"},
	}
	_, err := utils.C_books.UpdateOne(utils.Ctx, bson.D{{"_id", book.ID}}, data)
	if err != nil {
		fmt.Println("C_books.UpdateOne error:", err)
		return
	}
}
