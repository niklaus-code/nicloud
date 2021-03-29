package db

import (
  "fmt"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql" //这个一定要引入哦！！
 // "github.com/jinzhu/schema"
)

func Db() *gorm.DB {
  db, errDb:=gorm.Open("mysql","modis:modis@(10.0.90.151:3306)/blog")
  if errDb != nil {
    fmt.Println("11111111111111")
    fmt.Println(errDb)
  }

  return db
}
