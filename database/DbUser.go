package database

import (
	"Guestbook/models"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"gorm.io/gorm"
)

var db *sql.DB
var dB *gorm.DB

//连接数据库
func DbInit(){
	var err error
	db,err =sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/guestbook?charset=utf8&parseTime=true")
	fmt.Print(err)
	if err != nil {
		panic(fmt.Errorf("connect database Guestbook error\n"))
	}

}


func GetDb()*sql.DB{
	return db
}


//注册用户
func RegisterDb(user models.User) error {
	//判断去掉前后空格后是否为空
	if strings.TrimSpace(user.Password)== "" {
		return errors.New("密码为空!")
	}

	flag,err:=isUsernameExist(user.Username)
	if err != nil||!flag {
		return errors.New("用户名重复")
	}
	_, err = db.Exec(
		"insert into user(username,password,power) values(?,?,?)", user.Username, user.Password,user.Power)
	if err != nil {
		return errors.New("插入数据库失败")
	}
	return nil

}


//检查用户名是否存在
func isUsernameExist(username string) (bool, error) {
	stmt,err:=db.Query("select username,password from user;")
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	for stmt.Next() {
		var name,passwd  string
		err:=stmt.Scan(&name,&passwd)
		if err!= nil {
			return false, err
		}
		if username== name {
			return false, errors.New("用户名已经存在")
		}
	}
	return true,nil
}

//检查登陆是否正确
func IsUserTrue(username, password string) (bool, error,int) {
	rows,err:=db.Query("SELECT power from user where username=? and password=?",username,password)
	if err != nil {
		return false, errors.New("查找比对数据库失败"),-1
	}
	count:=0
	var power int
	for rows.Next() {
		rows.Scan(&power)
		count++
	}
	if count == 1 {
		return true,nil,power
	}else {
		return false, errors.New("密码或账户名错误"),-1
	}
}