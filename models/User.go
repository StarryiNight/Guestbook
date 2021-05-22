package models

type User struct {
	//用户 :账户名,密码
	Username string`form:"username" json:"username" ShouldBind:"required"`
	Password string`form:"password" json:"password" ShouldBind:"required"`
	Power int `form:"power" json:"power" ShouldBind:"required"`
}