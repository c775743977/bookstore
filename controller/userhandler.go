package controller

import (
	"bookstore1.1/dao"
	"fmt"
	"net/http"
	"html/template"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {  //处理登录界面
	if r.Header.Get("Cookie") != "" { //如果已有cookie说明已经登录，直接转主页
		Handler(w, r)
		return
	}
	username := r.PostFormValue("username")
		// 防止提交空白数据
		if username == "" {
			t := template.Must(template.ParseFiles("views/pages/user/login.html"))
			err := t.Execute(w, "")
			if err != nil {
				fmt.Println("t.Execute error:", err)
			}
			return
		}
	password := r.PostFormValue("password")
	// fmt.Println("POST数据：", username, password)
	uuid := dao.CheckUserNameAndPassword(username, password) //创造sessionID//判断账号密码
	if uuid != "" { 
		cookie := http.Cookie{
			Name : "user",
			Value : uuid,
			HttpOnly : true,
		}
		http.SetCookie(w, &cookie)  //创建cookie
		t := template.Must(template.ParseFiles("views/pages/user/login_success.html")) //正确转跳成功页面
		err := t.Execute(w, username)
		if err != nil {
			fmt.Println("t.Execute error:", err)
		}
	} else {
		t := template.Must(template.ParseFiles("views/pages/user/login.html")) //错误返回登录页面
		err := t.Execute(w, "用户名或密码错误")
		if err != nil {
			fmt.Println("t.Execute error:", err)
		}
	}
}

func ResgisterHandler(w http.ResponseWriter, r *http.Request) { //处理注册界面
	name := r.PostFormValue("username")
	// 防止提交空白数据
	if name == "" {
		t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
		err := t.Execute(w, "")
		if err != nil {
			fmt.Println("t.Execute error:", err)
		}
		return
	}
	passwd := r.PostFormValue("password")
	repwd := r.PostFormValue("repwd")
	email := r.PostFormValue("email")
	if passwd != repwd { //判断两次密码是否一致
		t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
		err := t.Execute(w, "两次密码不一致")
		if err != nil {
			fmt.Println("t.Execute error:", err)
		}
		return
	}
	if dao.CheckUserName(name) { //判断用户名是否存在
		dao.UserRegister(name, passwd, email) //不存在则注册用户
		dao.CreateCart(name)
		t := template.Must(template.ParseFiles("views/pages/user/regist_success.html")) //转跳注册成功页面
		err := t.Execute(w, "")
		if err != nil {
			fmt.Println("t.Execute error:", err)
		}
	} else {
		t := template.Must(template.ParseFiles("views/pages/user/regist.html")) //失败返回注册页面
		err := t.Execute(w, "用户名已存在")
		if err != nil {
			fmt.Println("t.Execute error:", err)
		}
	}
}


func LogOut(w http.ResponseWriter, r *http.Request) {  //注销
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		dao.DelSession(cookie.Value)
		cookie.MaxAge = -1 //立刻删除cookie
		// w.Header().Set("Cookie", cookie.String())
		http.SetCookie(w, cookie)
	}
	Handler(w, r)
}
