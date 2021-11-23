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

func MachineDb() (*gorm.DB,error) {
  db, err:=gorm.Open("mysql","modis:modis@(10.0.90.151:3306)/bigdata_machine?parseTime=true")
  db.SingularTable(true)
  if err != nil {
    return nil, err
  }

  sqlDB := db.DB()
  sqlDB.SetMaxIdleConns(10) //空闲连接数
  sqlDB.SetMaxOpenConns(100)//最大连接数

  return db, err
}

func NicloudDb() (*gorm.DB,error) {
  db, err:=gorm.Open("mysql","modis:modis@(127.0.0.1:3306)/nicloud?parseTime=true&loc=Local")
  if err != nil {
    return nil, err
  }

  sqlDB := db.DB()
  sqlDB.SetMaxIdleConns(100) //空闲连接数
  sqlDB.SetMaxOpenConns(1000)//最大连接数

  return db, err
}
