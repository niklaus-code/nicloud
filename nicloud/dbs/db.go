package db

import (
  "fmt"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  c "nicloud/config"
  "time"
)
var (
  config, _ = c.Exportconfig()
  nicloud = config.Nicloudb
  serveroom = config.Machinedb
)

func MachineDb() (*gorm.DB,error) {
  db, err:=gorm.Open("mysql",fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true"), serveroom.User, serveroom.Passwd, serveroom.Host, serveroom.Port, serveroom.Dbname)
  db.SingularTable(true)
  if err != nil {
    return nil, err
  }

  sqlDB := db.DB()
  sqlDB.SetMaxIdleConns(10000) //空闲连接数
  sqlDB.SetMaxOpenConns(10000)//最大连接数

  return db, err
}

func NicloudDb() (*gorm.DB,error) {
  db, err:=gorm.Open("mysql",fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true&loc=Local", nicloud.User, nicloud.Passwd, nicloud.Host, nicloud.Port, nicloud.Dbname))
  if err != nil {
    return nil, err
  }

  sqlDB := db.DB()
  sqlDB.SetMaxIdleConns(100) //空闲连接数
  sqlDB.SetMaxOpenConns(1000)//最大连接数
  sqlDB.SetConnMaxLifetime(time.Second * 6)

  return db, err
}
