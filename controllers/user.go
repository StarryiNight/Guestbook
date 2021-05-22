package controllers

import (
	"Guestbook/models"
	"github.com/gin-gonic/gin"
	"net/http"
)
import "Guestbook/database"

//注册
func Register(c *gin.Context)  {
	var user models.User
	user.Username=c.PostForm("username")
	user.Password=c.PostForm("password")
	key:=c.PostForm("key")
	if key=="Admin123456" {
		user.Power=1
	}else{
		user.Power= 0
	}

	err:=database.RegisterDb(user)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"ErrorTitle":"注册失败",
			"ErrorMessage":err.Error(),
		})
	}else {
		c.JSON(http.StatusOK,gin.H{
			"status:":http.StatusOK,
			"message":"注册成功",
		})
	}

}

//登陆
func Login(c *gin.Context) {
	username:=c.PostForm("username")
	password:=c.PostForm("password")
	flag,err,power:=database.IsUserTrue(username,password)
	var user string
	if power> 0 {
		user="administrator"
	}else {
		user="user"
	}
	if err != nil&&!flag {
		c.JSON(http.StatusBadRequest,gin.H{
			"ErrorTitle": "登陆失败",
			"ErrorMessage":err.Error(),
		})
	}else {
		c.SetCookie(
			"user",
			user,
			60,
			"/",
			"localhost",
			false,
			true,
		)
		c.JSON(http.StatusOK,gin.H{
			"status":http.StatusOK,
			"message":"登陆成功",
			"user":username,
		})
	}
}


//登出
func Logout(c *gin.Context)  {
	c.SetCookie(
		"user",
		"false",
		-1,
		"/",
		"localhost",
		false,
		true,)

	c.JSON(http.StatusMovedPermanently,"注销成功")
	c.Redirect(http.StatusMovedPermanently, "/login")

}