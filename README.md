## 留言板

##### 路由组

| /user |           |          |
| ----- | :-------: | -------- |
| POST  | /register | 用户注册 |
| POST  |  /login   | 用户登陆 |
| GET   |  /logout  | 用户登出 |

| /article |                  |                          |
| -------- | ---------------- | ------------------------ |
| /POST    | /public          | 发表评论                 |
| POST     | /publicReply     | 回复评论                 |
| GET      | /scanAll         | 查看所有评论             |
| GET      | /scanId          | 查看具体某ID评论         |
| PATCH    | /scanId/like     | 点赞                     |
| GET      | /scanReply/time  | 查看回复(按回复时间排序) |
| GET      | /scanReply/likes | 查看回复(按点赞次数)     |
| DELETE   | /deleteMessage   | 管理员删除评论           |

##### 用户和文章

| Post    |           |
| ------- | --------- |
| Id      | int       |
| Title   | string    |
| Content | string    |
| Author  | string    |
| Pid     | int       |
| Likes   | int       |
| Time    | time.Time |

| User     |        |
| -------- | ------ |
| Username | string |
| Password | string |
| Power    | int    |

