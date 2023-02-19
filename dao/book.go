package dao

import (
	"bookstore1.4/model"
	"bookstore1.4/utils"
	"strconv"
	_"fmt"
)

func GetBooks()  []*model.Book {
	var books []*model.Book
	utils.DBrr.RoundRobin().Find(&books)
	return books
}

func GetPage(bpageno string) (page *model.Page) {  //获取一页的图书数据
	pageno, _ := strconv.ParseInt(bpageno, 10, 0) //由pageno是从url的请求数据里读取的，所以默认是string类型
	page = &model.Page{}
	utils.DBrr.RoundRobin().Model(&model.Book{}).Count(&page.TotalBooks)
	page.PageNo = pageno
	page.PageSize = 4
	if page.TotalBooks % page.PageSize == 0 {
		page.TotalPages = page.TotalBooks / page.PageSize
	} else {
		page.TotalPages = page.TotalBooks / page.PageSize + 1
	}

	var books []*model.Book
	utils.DBrr.RoundRobin().Limit(int(page.PageSize)).Offset(int((page.PageNo-1)*page.PageSize)).Find(&books)
	page.Books = books
	return page
}

func GetPageByPrice(bpageno string, bmax string, bmin string) (page *model.Page) { //按照价格搜索结果的一页数据
	max, _ := strconv.ParseFloat(bmax, 64)
	min, _ := strconv.ParseFloat(bmin, 64)
	pageno, _ := strconv.ParseInt(bpageno, 10, 0)
	page = &model.Page{}
	utils.DBrr.RoundRobin().Model(&model.Book{}).Where("price between ? and ?", min, max).Count(&page.TotalBooks)
	page.PageNo = pageno
	page.PageSize = 4
	if page.TotalBooks % page.PageSize == 0 {
		page.TotalPages = page.TotalBooks / page.PageSize
	} else {
		page.TotalPages = page.TotalBooks / page.PageSize + 1
	}
	var books []*model.Book
	utils.DBrr.RoundRobin().Limit(int(page.PageSize)).Offset(int((page.PageNo-1)*page.PageSize)).Where("price between ? and ?", min, max).Find(&books)
	page.Books = books
	page.MaxPrice = max
	page.MinPrice = min
	return page
} 

func AddBook(book *model.Book) { //添加图书
	utils.WDB.Create(book)
}

func CheckBook(title string, author string) bool { //检查图书是否存在
	var res []string
	utils.DBrr.RoundRobin().Model(&model.Book{}).Where("title = ?", title).Select("author").Find(&res)
	for _,k := range res {
		if k == author {
			return true
		}
	}
	return false
}

func DelBook(id int64) { //删除图书
	utils.WDB.Where("id = ?", id).Delete(&model.Book{})
}

func GetBook(id int)  *model.Book { //根据输入ID获取对应图书
	var book model.Book
	utils.DBrr.RoundRobin().Where("id = ?", id).Find(&book)
	return &book
}

func AlterBook(book *model.Book) {  //更新图书信息
	utils.WDB.Where("id = ?", book.ID).Save(book)
}
