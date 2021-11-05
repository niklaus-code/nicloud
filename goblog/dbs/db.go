package db

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

func Db() (*gorm.DB,error) {
  db, errDb:=gorm.Open("mysql","modis:modis@(127.0.0.1:3306)/myblog?parseTime=true")
  defer db.Close()
  if errDb != nil {
    return nil, errDb
  }

  return db, nil
}

func NicloudDb() (*gorm.DB,error) {
  db, err:=gorm.Open("mysql","modis:modis@(127.0.0.1:3306)/nicloud?parseTime=true")
  if err != nil {
    return nil, err
  }

  return db, err
}
