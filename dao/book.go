package dao

import (
	"bookstore1.1/model"
	"bookstore1.1/utils"
	"fmt"
	"strconv"
)

func GetBooks() (books []*model.Book) {  //获取所有图书
	sqlstr := "select * from bookstore.books"
	rows, err := utils.DB.Query(sqlstr)
	if err != nil {
		fmt.Println("utils.DB.Query error:", err)
		return nil
	}
	for rows.Next() {
		book := &model.Book{}
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.Img_path)
		if err != nil {
			fmt.Println("rows.Scan error:", err)
		return nil
		}
		books = append(books, book)
	}
	return books
}

func AddBook(book *model.Book) { //添加图书
	sqlstr := "insert into bookstore.books(title, author, price, sales, stock, img_path) value(?,?,?,?,?,?)"
	_, err := utils.DB.Exec(sqlstr, book.Title, book.Author, book.Price, book.Sales, book.Stock, book.Img_path)
	if err != nil {
		fmt.Println("utils.DB.Exec error:", err)
		return
	}
}

func CheckBook(title string, author string) bool { //检查图书是否存在
	sqlstr := "select author from bookstore.books where title=?"
	row := utils.DB.QueryRow(sqlstr, title)
	var res string
	err := row.Scan(&res)
	if err != nil {
		fmt.Println("row.Scan error:", err)
		return false
	}
	if res == author {
		return true
	} else {
		return false
	}
}

func DelBook(id int64) { //删除图书
	sqlstr := "delete from bookstore.books where id=?"
	_, err := utils.DB.Exec(sqlstr, id)
	if err != nil {
		fmt.Println("utils.DB.Exec error:", err)
		return
	}
}

func GetBook(id int) ( *model.Book) { //根据输入ID获取对应图书
	sqlstr := "select * from bookstore.books where id=?"
	row := utils.DB.QueryRow(sqlstr, id)
	book := &model.Book{}
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.Img_path)
	if err != nil {
		fmt.Println("row.Scan error:", err)
	}
	return book
}

func AlterBook(book *model.Book) {  //更新图书信息
	sqlstr := "update bookstore.books set title=?, author=?, price=?, sales=?, stock=? where id=?"
	_, err := utils.DB.Exec(sqlstr, book.Title, book.Author, book.Price, book.Sales, book.Stock, book.ID)
	if err != nil {
		fmt.Println("utils.DB.Exec error:", err)
	}
}

func GetPage(bpageno string) (page *model.Page) {  //获取一页的图书数据
	pageno, _ := strconv.ParseInt(bpageno, 10, 0) //由pageno是从url的请求数据里读取的，所以默认是string类型
	sqlstr := "select count(*) from bookstore.books"
	row := utils.DB.QueryRow(sqlstr)
	page = &model.Page{}
	err := row.Scan(&page.TotalBooks)
	if err != nil {
		fmt.Println("row.Scan GetPage error:", err)
		return nil
	}
	page.PageNo = pageno
	page.PageSize = 4
	if page.TotalBooks % page.PageSize == 0 {
		page.TotalPages = page.TotalBooks / page.PageSize
	} else {
		page.TotalPages = page.TotalBooks / page.PageSize + 1
	}
	sqlstr = "select * from bookstore.books limit ?,?"
	rows, err := utils.DB.Query(sqlstr, (page.PageNo - 1) * page.PageSize, page.PageSize)
	if err != nil {
		fmt.Println("utils.DB.Query GetPage error:", err)
		return nil
	}
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.Img_path)
		if err != nil {
			fmt.Println("rows.Scan error:", err)
		return nil
		}
		books = append(books, book)
	}
	page.Books = books
	return page
}

func GetPageByPrice(bpageno string, bmax string, bmin string) (page *model.Page) { //按照价格搜索结果的一页数据
	max, _ := strconv.ParseFloat(bmax, 64)
	min, _ := strconv.ParseFloat(bmin, 64)
	pageno, _ := strconv.ParseInt(bpageno, 10, 0)
	sqlstr := "select count(*) from bookstore.books where price between ? and ?"
	row := utils.DB.QueryRow(sqlstr, min, max)
	page = &model.Page{}
	err := row.Scan(&page.TotalBooks)
	if err != nil {
		fmt.Println("row.Scan GetPage error:", err)
		return nil
	}
	page.PageNo = pageno
	page.PageSize = 4
	if page.TotalBooks % page.PageSize == 0 {
		page.TotalPages = page.TotalBooks / page.PageSize
	} else {
		page.TotalPages = page.TotalBooks / page.PageSize + 1
	}
	sqlstr = "select * from bookstore.books where price between ? and ? limit ?,?"
	rows, err := utils.DB.Query(sqlstr, min, max, (page.PageNo - 1) * page.PageSize, page.PageSize)
	if err != nil {
		fmt.Println("utils.DB.Query GetPage error:", err)
		return nil
	}
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.Img_path)
		if err != nil {
			fmt.Println("rows.Scan error:", err)
		return nil
		}
		books = append(books, book)
	}
	page.Books = books
	page.MaxPrice = max
	page.MinPrice = min
	return page
} 
