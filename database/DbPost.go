package database

import (
	"Guestbook/models"
	"fmt"
)

//发表回复或评论 (评论 id=pid,回复 pid=回复id)
func DbPublic(post models.Post) (bool, error)  {
	sql1:="insert into post(title,content,author,id,likes,pid,time)values(?,?,?,?,?,?,?)"
	fmt.Println(post)
	_, err := db.Exec(sql1,post.Title, post.Content, post.Author, post.Id, post.Likes, post.Pid,post.Time)
	if err != nil {
		return false, err
	}
	return true,nil

}

//浏览所有评论(不包括回复)
func DbScanAll() ([]models.Post ,error){
	rows, err:= db.Query("select * from post where pid=Id")
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		post:=models.Post{}
		if e:=rows.Scan(&post.Id,&post.Title,&post.Content,&post.Author,&post.Likes,&post.Pid,&post.Time);e!= nil {
			fmt.Println(e)
			return nil, e
		}
		posts=append(posts,post)
	}
	return posts,nil
}

//以id浏览
func DbScanId(id int) (models.Post, error) {
	var post models.Post
	row :=db.QueryRow("select * from post where id = ? ", id)
	if e:=row.Scan(&post.Id,&post.Title,&post.Content,&post.Author,&post.Likes,&post.Pid,&post.Time);e!= nil {
		return post, e
	}
	return post,nil
}

//浏览回复 时间排序
func DbscanReplyByTime(id int)([]models.Post , error)  {
	rows, err:= db.Query("select * from post WHERE pid=? ORDER BY time ASC",id)
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		post:=models.Post{}
		if e:=rows.Scan(&post.Id,&post.Title,&post.Content,&post.Author,&post.Likes,&post.Pid,&post.Time);e!= nil {
			fmt.Println(e)
			return nil, e
		}
		posts=append(posts,post)
	}
	return posts,nil

}

//浏览回复 点赞数排序
func DbscanReplyByLikes(id int)([]models.Post , error)  {
	rows, err:= db.Query("select * from post WHERE pid=? ORDER BY likes ASC",id)
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		post:=models.Post{}
		if e:=rows.Scan(&post.Id,&post.Title,&post.Content,&post.Author,&post.Likes,&post.Pid,&post.Time);e!= nil {
			fmt.Println(e)
			return nil, e
		}
		posts=append(posts,post)
	}
	return posts,nil

}

//点赞
func DbLike(id int)(error) {
	_,err:=db.Query("UPDATE post set likes=likes+1  WHERE Id= ?", id)
	if err != nil {
		return err
	}
	return nil
}

func DbDeleteMessage(id int)(error) {
	_,err:= db.Exec("delete from post where id=?", id)
	if err != nil {
		return err
	}else{
		return nil
	}
}