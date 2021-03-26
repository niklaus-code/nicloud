package common

import "github.com/niklaus-code/gogogo/db"

//定义结构体 #类似初始化dict
type Blog struct {
    Id string `json:"id"` //结构体变量首字母大，小写字母为私有变量，c.JSON()调用JSON/marshal()进行序列化, 属于包外方法, 无法解析到小写字母开头的成员变量
    Title string `json:"title"`
    Content string `json:"content"`
    Create_time string `json:"create_time"`
    Img string `json:"img"`
    Category_name string `json:"category_name"`
    }

func BlogGet(startpage int, offset int) ([]*Blog) {
    //初始化db
    db := db.Db()

    //查询
    rows, err := db.Query("select m.id,m.title,m.content, m.img, group_concat(i.category_name) from myblog_list m LEFT JOIN article_class a on m.id = a.article_id  LEFT JOIN myblog_caegory_info i on a.category_id = i.category_id where m.status = 1 GROUP BY m.id order by m.id desc limit ?, ?", startpage, offset)

    if err != nil {
        return nil
        }

    //定义struct数组
    var bloglist []*Blog

    //循环遍历添加到struct数组
    for rows.Next(){
        var blog Blog
        rows.Scan(&blog.Id, &blog.Title, &blog.Content, &blog.Img, &blog.Category_name)
        bloglist = append(bloglist, &blog)
        }

    //关闭数据库连接
    rows.Close()
    return bloglist
    }


type Thoughts struct {
    Content string `json:"content"`
    Create_time string `json:"create_time"`
    Like_number int `json:"like_number"`
    }


func ThoughtsGet() ([]*Thoughts) {
    //初始化db
    db := db.Db()

    //查询
    rows, err := db.Query("SELECT content,create_time,like_number from myblog_thoughts order by id desc")

    if err != nil {
        return nil
        }

    //定义struct数组
    var thoughtslist []*Thoughts

    //循环遍历添加到struct数组
    for rows.Next(){
        var thoughts Thoughts
        rows.Scan(&thoughts.Content, &thoughts.Create_time, &thoughts.Like_number)
        thoughtslist = append(thoughtslist, &thoughts)
        }

    //关闭数据库连接
    rows.Close()
    return thoughtslist
    }


type Read struct {
    Id int `json:"id"`
    Title string `json:"title"`
    Beginning string `json:"beginning"`
    Author string `json:"author"`
    Image string `json:"image"`
}

func ReadGet() ([]*Read) {
    //初始化db
    db := db.Db()

    //查询
    rows, err := db.Query("select id, title, beginning, author, image from myblog_read where status=1 order by create_time desc")

    if err != nil {
        return nil
        }

    //定义struct数组
    var readlist []*Read

    //循环遍历添加到struct数组
    for rows.Next(){
        var read Read
        rows.Scan(&read.Id, &read.Title, &read.Beginning, &read.Author, &read.Image)
        readlist = append(readlist, &read)
        }

    //关闭数据库连接
    rows.Close()
    return readlist
}


func BlogGetById(id int) (*Blog) {
    //初始化db
    db := db.Db()

    //查询
    rows, err := db.Query("select t1.id, t1.title,t1.content, t1.create_time, t3.category_name from myblog_list t1 left join article_class t2 on t1.id=t2.article_id LEFT JOIN myblog_caegory_info t3 on t2.category_id = t3.category_id   where t1.id = ?;", id)

    if err != nil {
        return nil
        }


    var blog Blog
    //循环遍历添加到struct数组
    for rows.Next(){
        rows.Scan(&blog.Id, &blog.Title, &blog.Content, &blog.Create_time, &blog.Category_name)
        }

    //关闭数据库连接
    rows.Close()
    return &blog
}
