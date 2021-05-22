package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//检查是否登陆
func LoginCheck() gin.HandlerFunc {
	return func(c *gin.Context) {

		res,err:=c.Cookie("user")
		if err != nil ||res==""{
			c.JSON(http.StatusUnauthorized, "请先登陆")
			c.Redirect(http.StatusMovedPermanently, "/login")
			c.Abort()
		}
	}
}
