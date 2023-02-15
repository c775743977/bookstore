package controller

import (
	"github.com/gin-gonic/gin"
	"bookstore1.4/dao"
	_"bookstore1.4/model"
	"net/http"
	"strconv"
	_"fmt"
)

func ShowCart(c *gin.Context) {
	cookie, _ := c.Cookie("user")
	sess := dao.GetSession(cookie)
	cart := dao.GetCart(sess.UserID)
	cart.UserName = sess.Username
	c.HTML(http.StatusOK, "cart/cart.html", cart)
}

func AddBookToCart(c *gin.Context) {
	bookid, _ := strconv.ParseInt(c.PostForm("bookId"), 0, 0)
	cookie, err := c.Cookie("user")
	if err != nil {
		c.String(400, "请先登录...")
		return
	}
	sess := dao.GetSession(cookie)
	cart := dao.GetCart(sess.UserID)
	dao.AddItem(int(bookid), cart.ID)
	dao.UpdateCart(sess.UserID)
	Handler(c)
}

func DelItemHandler(c *gin.Context) {
	cookie, _ := c.Cookie("user")
	sess := dao.GetSession(cookie)
	cart := dao.GetCart(sess.UserID)
	bookid, _ := strconv.ParseInt(c.Query("bookID"), 0, 0)
	dao.DelItem(int(bookid), cart.ID)
	dao.UpdateCart(sess.UserID)
	cart = dao.GetCart(sess.UserID)
	cart.UserName = sess.Username
	c.HTML(http.StatusOK, "cart/cart.html", cart)
}

func CleanCartHandler(c *gin.Context) {
	cookie, _ := c.Cookie("user")
	sess := dao.GetSession(cookie)
	dao.CleanCart(sess.UserID)
	cart := dao.GetCart(sess.UserID)
	cart.UserName = sess.Username
	c.HTML(http.StatusOK, "cart/cart.html", cart)
}

func CartAlter(c *gin.Context) {
	cookie, _ := c.Cookie("user")
	booknum, _ := strconv.ParseInt(c.PostForm("BookNum"), 0, 0)
	bookid, _ := strconv.ParseInt(c.Query("BookID"), 0, 0)
	sess := dao.GetSession(cookie)
	cart := dao.GetCart(sess.UserID)
	dao.ModifyNum(int(bookid), cart.ID, int(booknum))
	dao.UpdateCart(sess.UserID)
	cart = dao.GetCart(sess.UserID)
	cart.UserName = sess.Username
	ShowCart(c)
}