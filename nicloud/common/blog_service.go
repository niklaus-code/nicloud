package common

import (
	"fmt"
	"goblog/dbs"
)

//定义结构体 #类似初始化dict
type Blog struct {
	Id            string `json:"id"` //结构体变量首字母大，小写字母为私有变量，c.JSON()调用JSON/marshal()进行序列化, 属于包外方法, 无法解析到小写字母开头的成员变量
	Title         string `json:"title"`
	Content       string `json:"content"`
	Create_time   string `json:"create_time"`
	Category_name string `json:"category_name"`
}

func BlogGet(startpage int, offset int) ([]*Blog, error) {
	//初始化db
	db, err := db.Db()
	if err != nil {
		return nil, err
	}
	row := db.Order("create_time desc").Table("myblog_lists m").Select("m.id, m.title, m.create_time, m.content, myblog_caegory_info.category_name").Joins("LEFT JOIN article_class on m.id = article_class.article_id").Joins("left join myblog_caegory_info on myblog_caegory_info.category_id = article_class.category_id where m.status=1").Limit(offset).Offset((startpage - 1) * offset)
	/*
	   sql := fmt.Sprintf("select id, title, content,  img, group_concat(i.category_name) from myblog_lists as m" +
	     "LEFT JOIN article_class a on m.id = a.article_id " +
	     "LEFT JOIN myblog_caegory_info i on a.category_id = i.category_id " +
	     "where m.status = 1 GROUP BY m.id order by m.id desc limit %d, %d", startpage, offset)
	*/
	var blog []*Blog
	row.Scan(&blog)
	return blog, nil
}

type Myblog_thoughts struct {
	Content     string `json:"content"`
	Create_time string `json:"create_time"`
	Like_number int    `json:"like_number"`
}

func ThoughtsGet() ([]*Myblog_thoughts, error) {
	//初始化db
	db, err := db.Db()
	if err != nil {
		return nil, err
	}
	var myblog_thoughts []*Myblog_thoughts
	//查询

	db.Order("id desc").Find(&myblog_thoughts)
	return myblog_thoughts, nil
}

type Myblog_reads struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Beginning string `json:"beginning"`
	Author    string `json:"author"`
	Image     string `json:"image"`
}

func ReadGet() ([]*Myblog_reads, error) {
	//初始化db
	db, err := db.Db()
	if err != nil {
		return nil, err
	}
	//查询
	res := db.Model(&Myblog_reads{})
	if res != nil {
		fmt.Println(res)
	}
	var blog []*Myblog_reads
	err1 := res.Find(&blog)
	if err1.Error != nil {
		return nil, err1.Error
	}
	return blog, nil
}

type Bloginfo struct {
	Id            int
	Title         string
	Content       string
	Create_time   string
	Category_name string
}

func BlogGetById(id int) (Bloginfo, error) {
	//初始化db
	db, err := db.Db()
	if err != nil {
		return Bloginfo{}, err
	}
	var bloginfo Bloginfo
	//*查询
	row := db.Table("myblog_lists m").Select("m.id,  m.title, m.content, m.create_time, myblog_caegory_info.category_name").Joins("left join article_class on article_class.article_id = m.id ").Joins("left join myblog_caegory_info on article_class.category_id = myblog_caegory_info.category_id").Where(fmt.Sprintf("m.id=%d", id))
	/*
	   rows := fmt.Sprint("select t1.id, t1.title,t1.content, t1.create_time, t3.category_name from myblog_list t1" +
	     " left join article_class t2 on t1.id=t2.article_id " +
	     "LEFT JOIN myblog_caegory_info t3 on t2.category_id = t3.category_id " +
	     "  where t1.id = %d;", id)
	*/
	row.Scan(&bloginfo)

	return bloginfo, nil
}
