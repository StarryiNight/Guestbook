package models

import "time"

type Post struct {
	Id int
	Title string
	Content string
	Author string
	Pid int
	Likes int
	Time time.Time
}

