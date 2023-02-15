package controller

import (
	"github.com/gin-gonic/gin"
	"bookstore1.4/dao"
	"bookstore1.4/model"
	"net/http"
	_"fmt"
	"strconv"
)

//处理登录请求
func LoginHandler(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	uuid := dao.CheckUserNameAndPassword(username, password)
	if uuid != "" {
		c.SetCookie("user", uuid, 0, "/", "localhost", true, true)
		c.HTML(http.StatusOK, "user/login_success.html", username)
	} else {
		c.HTML(400, "user/login.html", "用户名或密码错误")
	}
}

func RegistHandler(c *gin.Context) { //处理注册请求
	user := &model.User{Privilege : "N",}
	err := c.Bind(&user)
	if err != nil {
		c.HTML(400, "user/regist.html", "用户名和密码不能为空")
		return
	}
	if dao.CheckUserName(user.Name) {
		if repwd := c.PostForm("repwd"); repwd == user.Password {
			dao.UserRegist(user)
			dao.CreateCart(user.Name)
			c.HTML(200, "user/regist_success.html", user.Name)
			return
		} else {
			c.HTML(400, "user/regist.html", "两次密码不一致")
			return
		}
	} else {
		c.HTML(400, "user/regist.html", "该用户名已存在")
		return
	}
}

func UsersManageHandler(c *gin.Context) { //显示所有用户信息
	uuid, _ := c.Cookie("user")
	sess := dao.GetSession(uuid)
	if sess.Username == "root" {
		us := dao.GetAllUsers()
		c.HTML(200, "manager/user_manager.html", us)
	} else {
		c.String(400, "您没有权限访问此页面！")
	}
}

func LogoutHandler(c *gin.Context) {
	uuid, _ := c.Cookie("user")
	dao.DelSession(uuid)
	c.SetCookie("user", uuid, -1, "/", "localhost", true, true)
	Handler(c)
}

func ToAlterUserHandler(c *gin.Context) { //前往更新用户信息页面
	id, _ := strconv.ParseInt(c.Query("userID"), 10, 0)
	if id > 0 {
		user := dao.GetUserByID(int(id))
		c.HTML(200, "manager/user_edit.html", user)
	} else {
		c.HTML(200, "manager/user_edit.html", nil)
	}
}

func AlterUserHandler(c *gin.Context) { //更改用户信息或者添加用户
	var user model.User
	userid := c.Query("userID")
	id, _ := strconv.ParseInt(userid, 10, 0)
	u := dao.GetUserByID(int(id))
	err := c.Bind(&user)
	//报错是说明提交了空白数据
	if err != nil {
		user.Err = "失败！不允许提交空白数据"
		c.HTML(400, "manager/user_edit.html", user)
		return
	}
	if user.ID == 0 {
		if !dao.CheckUserName(user.Name) { //判断图书是否存在
			user.Err = "失败！该用户已存在"
			c.HTML(400, "manager/user_edit.html", user)
			return
		}
		dao.UserRegist(&user)
		UsersManageHandler(c)
	} else {
		if u.Name != user.Name && !dao.CheckUserName(user.Name) {
			user.Err = "失败！该用户已存在"
			c.HTML(400, "manager/user_edit.html", user)
			return
		}
		dao.AlterUser(&user)
		UsersManageHandler(c)
	}
}

func DelUserHandler(c *gin.Context) {  //删除用户
	id := c.Query("userID")
	if dao.DelUser(id) {
		UsersManageHandler(c)
	} else {
		c.String(400, "无法删除用户，该用户存在未完成的订单")
	}
}