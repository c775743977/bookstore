package controller

import (
	"Hello_golang/bookstore/dao"
	"Hello_golang/bookstore/model"
	"fmt"
	"net/http"
	"html/template"
	"strconv"
)

func Handler(w http.ResponseWriter, r *http.Request) { //主页
	// t := template.Must(template.ParseFiles("views/index.html"))
	pageno := r.FormValue("PageNo")
	min := r.FormValue("min")
	max := r.FormValue("max")
	cookie, _ := r.Cookie("user")
	username := ""
	if cookie != nil {
		username, _ = dao.GetSession(cookie.Value)
	}
	// fmt.Println("cookie=", cookie.String())
	if (min == "" && max == "") || (min == "0" && max == "0"){
		if pageno == "" {
			pageno = "1"
		}
		page := dao.GetPage(pageno)
		page.Username = username
		t := template.Must(template.ParseFiles("views/index.html"))
		// fmt.Println("page.username=", page.Username)
		t.Execute(w, page)
	} else {
		if pageno == "" {
			pageno = "1"
		}
		page := dao.GetPageByPrice(pageno, max, min)
		page.Username = username
		t := template.Must(template.ParseFiles("views/index.html"))
		t.Execute(w, page)
	}
}

func GetBooksHandler(w http.ResponseWriter, r *http.Request) { //获取所有图书
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	books := dao.GetBooks()
	t.Execute(w, books)
}

// func AddBookHandler(w http.ResponseWriter, r *http.Request) { //添加新图书
// 	var book model.Book
// 	book.Title = r.PostFormValue("title")
// 	fmt.Println("传入参数:", book.Title)
// 	if book.Title == "" { //防止传入空数据
// 		t := template.Must(template.ParseFiles("views/pages/manager/book_alter.html"))
// 		t.Execute(w, "")
// 		return
// 	}
// 	book.Price, _ = strconv.ParseFloat(r.PostFormValue("price"), 64)
// 	book.Author = r.PostFormValue("author")
// 	book.Sales, _ = strconv.ParseInt(r.PostFormValue("sales"), 10, 0)
// 	book.Stock, _ = strconv.ParseInt(r.PostFormValue("stock"), 10, 0)
// 	book.Img_path = "static/img/default.jpg"
// 	if dao.CheckBook(book.Title, book.Author) { //判断图书是否存在
// 		t := template.Must(template.ParseFiles("views/pages/manager/book_alter.html"))
// 		t.Execute(w, "该书已存在!")
// 	} else {
// 		dao.AddBook(&book)
// 		t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html")) //添加成功
// 		books := dao.GetBooks()
// 		t.Execute(w, books)
// 	}
// 	fmt.Println(book)
// }


func DelBookHandler(w http.ResponseWriter, r *http.Request) { //删除图书
	id, _ := strconv.ParseInt(r.FormValue("bookID"), 10, 0) //从URL中获取请求数据
	dao.DelBook(id)
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	books := dao.GetBooks()
	t.Execute(w, books)
}

func GetBookHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(r.FormValue("bookID"), 10, 0)
	book := dao.GetBook(int(id))
	t := template.Must(template.ParseFiles("views/pages/manager/book_alter.html"))
	t.Execute(w, book)
}

// func AlterBookHandler(w http.ResponseWriter, r *http.Request) {
// 	id, _ := strconv.ParseInt(r.PostFormValue("bookID"), 10, 0)
// 	book := dao.GetBook(int(id))
// 	book.Title = r.PostFormValue("title")
// 	book.Author = r.PostFormValue("author")
// 	book.Price, _ = strconv.ParseFloat(r.PostFormValue("price"), 64)
// 	book.Sales, _ = strconv.ParseInt(r.PostFormValue("sales"), 10, 0)
// 	book.Stock, _ = strconv.ParseInt(r.PostFormValue("stock"), 10, 0)
// 	fmt.Println("提交的数据:", book)
// 	if book.Title == "" {
// 		return
// 	}
// 	dao.AlterBook(book)
// 	GetBooksHandler(w, r)
// }

func ToAddOrAlterBook(w http.ResponseWriter, r *http.Request) { //根据bookID有无跳转添加或修改图书页面
	id, _ := strconv.ParseInt(r.FormValue("bookID"), 10, 0)
	if id > 0 {
		book := dao.GetBook(int(id))
		t := template.Must(template.ParseFiles("views/pages/manager/book_alter.html"))
		t.Execute(w, book)
	} else {
		t := template.Must(template.ParseFiles("views/pages/manager/book_alter.html"))
		t.Execute(w, "")
	}
}

func AddOrAlterBook(w http.ResponseWriter, r *http.Request) { //执行添加或者修改操作
	var book model.Book
	book.ID, _ = strconv.ParseInt(r.PostFormValue("bookID"), 10, 0)
	book.Title = r.PostFormValue("title")
	book.Price, _ = strconv.ParseFloat(r.PostFormValue("price"), 64)
	book.Author = r.PostFormValue("author")
	book.Sales, _ = strconv.ParseInt(r.PostFormValue("sales"), 10, 0)
	book.Stock, _ = strconv.ParseInt(r.PostFormValue("stock"), 10, 0)
	book.Img_path = "static/img/default.jpg"
	if book.ID == 0 {
		if book.Title == "" { //防止传入空数据
		t := template.Must(template.ParseFiles("views/pages/manager/book_alter.html"))
		t.Execute(w, "")
		return
		}
		if dao.CheckBook(book.Title, book.Author) { //判断图书是否存在
			t := template.Must(template.ParseFiles("views/pages/manager/book_alter.html"))
			t.Execute(w, book)
		} else {
			fmt.Println("ADD提交数据为:", book)
			dao.AddBook(&book)
			GetPageHandler(w, r)
		}
	} else {
		fmt.Println("ALTER提交数据为:", book)
		dao.AlterBook(&book)
		GetPageHandler(w, r)
	}
}

func GetPageHandler(w http.ResponseWriter, r *http.Request) { //对某一页初始化
	pageno := r.FormValue("PageNo")
	if pageno == "" {
		pageno = "1"
	}
	page := dao.GetPage(pageno)
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	t.Execute(w, page)
}

// func GetPageByPriceHandler(w http.ResponseWriter, r *http.Request) {
// 	pageno := r.FormValue("PageNo")
// 	max := r.PostFormValue("max")
// 	min := r.PostFormValue("min")
// 	if pageno == "" {
// 		pageno = "1"
// 	}
// 	page := dao.GetPageByPrice(pageno, max, min)
// 	fmt.Println("page=", page)
// 	t := template.Must(template.ParseFiles("views/index.html"))
// 	t.Execute(w, page)
// }