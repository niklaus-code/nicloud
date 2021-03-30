package db

import (
  "fmt"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql" //这个一定要引入哦！！
 // "github.com/jinzhu/schema"
)

func Db() *gorm.DB {
  db, errDb:=gorm.Open("mysql","ysman:123456@(127.0.0.1:3306)/myblog")
  if errDb != nil {
    fmt.Println(errDb)
  }

  return db
}
