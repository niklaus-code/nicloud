package db

import (
  "fmt"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

func Db() *gorm.DB {
  db, errDb:=gorm.Open("mysql","modis:modis@(127.0.0.1:3306)/myblog?parseTime=true")
  if errDb != nil {
    fmt.Println(errDb)
  }

  return db
}

func NicloudDb() (*gorm.DB,error) {
  db, err:=gorm.Open("mysql","modis:modis@(127.0.0.1:3306)/nicloud?parseTime=true")
  if err != nil {
    return nil, err
  }

  return db, err
}
