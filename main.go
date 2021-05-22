package main

import (
	"Guestbook/Routers"
	"Guestbook/database"
	"github.com/gin-gonic/gin"
)

func main() {
	database.DbInit()
	r:=gin.Default()
	Routers.Router(r)
	r.Run()
}
