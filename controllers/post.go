package controllers

import (
	"Guestbook/database"
	"Guestbook/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

//发表评论
func Public(c *gin.Context) {
	var post models.Post
	post.Title = c.PostForm("Title")
	post.Author = c.PostForm("Author")
	post.Content = c.PostForm("Content")
	post.Id, _ = strconv.Atoi(c.PostForm("Id"))
	post.Pid = post.Id
	post.Likes = 0
	now := time.Now()
	mm, _ := time.ParseDuration("8h")
	post.Time = now.Add(mm)

	flag, err := database.DbPublic(post)
	if err != nil && flag == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"ErrorTitle":   "发表评论失败",
			"ErrorMessage": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "发表评论成功",
		})
	}
}

//浏览所有评论 不包括评论下的回复
func ScanAll(c *gin.Context) {
	posts, err := database.DbScanAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
	}
	for _, i := range posts {
		c.JSON(http.StatusOK, gin.H{
			"Id":      i.Id,
			"Title":   i.Title,
			"Author":  i.Author,
			"Content": i.Content,
			"Likes":   i.Likes,
			"Time":    i.Time,
		})

	}
}

//以id浏览回复或留言
func ScanId(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("Id"))
	post, err := database.DbScanId(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Id":      post.Id,
			"Title":   post.Title,
			"Author":  post.Author,
			"Content": post.Content,
			"Likes":   post.Likes,
			"Time":    post.Time,
		})
	}
}

//查看留言 时间排序
func ScanReplyByTime(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("Id"))
	posts, err := database.DbscanReplyByTime(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
	}
	for _, i := range posts {
		c.JSON(http.StatusOK, gin.H{
			"Id":      i.Id,
			"Title":   i.Title,
			"Author":  i.Author,
			"Content": i.Content,
			"Likes":   i.Likes,
			"Time":    i.Time,
		})

	}
}

//查看留言 点赞数排序
func ScanReplyByLikes(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("Id"))
	posts, err := database.DbscanReplyByLikes(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
	}
	for _, i := range posts {
		c.JSON(http.StatusOK, gin.H{
			"Id":      i.Id,
			"Title":   i.Title,
			"Author":  i.Author,
			"Content": i.Content,
			"Likes":   i.Likes,
			"Time":    i.Time,
		})

	}
}

//点赞
func Like(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("Id"))
	err := database.DbLike(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "点赞成功",
		})
	}
}

//发表回复
func PublicReply(c *gin.Context) {
	var post models.Post
	post.Title = c.PostForm("Title")
	post.Author = c.PostForm("Author")
	post.Content = c.PostForm("Content")
	post.Id, _ = strconv.Atoi(c.PostForm("Id"))
	post.Pid, _ = strconv.Atoi(c.PostForm("pid"))
	post.Likes = 0
	now := time.Now()
	mm, _ := time.ParseDuration("8h")
	post.Time = now.Add(mm)

	flag, err := database.DbPublic(post)
	if err != nil && flag == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"ErrorTitle":   "回复失败",
			"ErrorMessage": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "回复成功",
		})
	}
}

func DeleteMessage(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("Id"))
	res, err := c.Cookie("user")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"ErrorTitle":   "请求失败",
			"ErrorMessage": err.Error(),
		})
	} else if res != "administrator" {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error":   "没有权限进行此操作",
		})
	} else {
		err := database.DbDeleteMessage(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"ErrorTitle":   "操作错误",
				"ErrorMessage": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":  http.StatusOK,
				"message": "删除成功",
			})
		}
	}
}
