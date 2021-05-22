package Routers

import (
	"Guestbook/controllers"
	"Guestbook/middleware"
	"github.com/gin-gonic/gin"
)


func Router(r *gin.Engine)  {

	user:=r.Group("/user")
	{
		//登陆
		user.POST("/login",controllers.Login)
		//注册
		user.POST("/register",controllers.Register)
		//登出
		user.GET("/logout",controllers.Logout)
	}
	article:=r.Group("/article")
	{
		//发表评论
		article.POST("/public",middleware.LoginCheck(),controllers.Public)
		//回复评论
		article.POST("/publicReply",middleware.LoginCheck(),controllers.PublicReply)
		//查看所有评论
		article.GET("/scanAll",middleware.LoginCheck(),controllers.ScanAll)
		//查看第id个评论
		article.GET("/scanId",middleware.LoginCheck(),controllers.ScanId)
		//点赞
		article.PATCH("/scanId/like",controllers.Like)
		//查看回复(按回复时间排序)
		article.GET("/scanReply/time",middleware.LoginCheck(),controllers.ScanReplyByTime)
		//查看回复(按点赞次数)
		article.GET("/scanReply/likes",middleware.LoginCheck(),controllers.ScanReplyByLikes)
		//管理员删除评论
		article.DELETE("/deleteMessage",middleware.LoginCheck(),controllers.DeleteMessage)
	}
	
}
