package controller

import (
	"github.com/gin-gonic/gin"
	"bookstore1.4/dao"
	"bookstore1.4/model"
	"net/http"
	"strconv"
	_"fmt"
)

//主页
func Handler(c *gin.Context) {  //首页
	page := &model.Page{}
	pageno := c.Query("PageNo")
	//因为第一次提交的价格数据是通过表单，然后被写进url里，所以应该先判断表单是否有数据再判断url是否有数据
	//如果按照标准库来写，FormValue可以读取url和表单的数据，谁有数据就用谁的数据，就不会需要再分别去读取
	//先从表单读取数据
	min := c.PostForm("min")
	max := c.PostForm("max")
	//如果表单没有数据再读取url的数据
	if min == "" && max == "" {
		min = c.Query("min")
		max = c.Query("max")
	}
	cookie, err := c.Cookie("user")
	sess := &model.Session{}
	//判断cookie是否存在，如果存在就获取登录的用户
	if err == nil {
		sess = dao.GetSession(cookie)
	}
	//判断max和min是否有值
	if (min == "" && max == "") || (min == "0" && max == "0"){
		//判断是否有页数，如果没有就显示第一页
		if pageno == "" {
			pageno = "1"
		}
		page = dao.GetPage(pageno)
	} else {
		if pageno == "" {
			pageno = "1"
		}
		page = dao.GetPageByPrice(pageno, max, min)
	}
	page.Username = sess.Username
	c.HTML(http.StatusOK, "index/index.html", page)
}

func GetBooksHandler(c *gin.Context) { //获取所有图书
	books := dao.GetBooks()
	c.HTML(200, "manager/book_manager.html", books)
}

func DelBookHandler(c *gin.Context) {
	//获取要删除的图书ID
	id, _ := strconv.ParseInt(c.Query("bookID"), 10, 0)
	dao.DelBook(id)
	//再次渲染删除后的那一页html
	GetPageHandler(c)
}

//对某项图书数据进行修改
func GetBookHandler(c *gin.Context) {
	//获取要改动的图书ID
	id, _ := strconv.ParseInt(c.Query("bookID"), 10, 0)
	book := dao.GetBook(int(id))
	c.HTML(200, "manager/book_edit.html", book)
}

//根据bookID有无跳转添加或修改图书页面,有ID就去修改，无ID就说明要添加
func ToAddOrAlterBook(c *gin.Context) { 
	id, _ := strconv.ParseInt(c.Query("bookID"), 10, 0)
	if id > 0 {
		book := dao.GetBook(int(id))
		c.HTML(200, "manager/book_edit.html", book)
	} else {
		c.HTML(200, "manager/book_edit.html", nil)
	}
}

func AddOrAlterBook(c *gin.Context) { //执行添加或者修改操作
	var book model.Book
	bookid := c.Query("bookID")
	id, _ := strconv.ParseInt(bookid, 10, 0)
	bbook := dao.GetBook(int(id))
	err := c.Bind(&book)
	book.Img_path = "static/img/default.jpg"
	//报错是说明提交了空白数据
	if err != nil {
		book.Err = "失败！不允许提交空白数据"
		c.HTML(400, "manager/book_edit.html", book)
		return
	}
	if book.ID == 0 {
		if dao.CheckBook(book.Title, book.Author) { //判断图书是否存在
			book.Err = "失败！该书已存在"
			c.HTML(400, "manager/book_edit.html", book)
			return
		}
		dao.AddBook(&book)
		GetPageHandler(c)
	} else {
		if bbook.Title != book.Title && dao.CheckBook(book.Title, book.Author) {
			book.Err = "失败！该书已存在"
			c.HTML(400, "manager/book_edit.html", book)
			return
		}
		dao.AlterBook(&book)
		GetPageHandler(c)
	}
}

func GetPageHandler(c *gin.Context) { //对某一页初始化,不带封面
	cookie, _ := c.Cookie("user")
	sess := dao.GetSession(cookie)
	// if sess.Username != "root" {
	// 	return
	// }
	pageno := c.Query("PageNo")
	if pageno == "" {
		pageno = "1"
	}
	page := dao.GetPage(pageno)
	page.Username = sess.Username
	c.HTML(200, "manager/book_manager.html", page)
}