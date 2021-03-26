package db

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "github.com/go-ini/ini"
    "fmt"
    "strings"
    )

func Db() *sql.DB {
    cfg, err := ini.Load("conf/setting.ini")
    if err != nil {
        fmt.Println(err)
    }

   var ip = cfg.Section("mysql").Key("ip").String()
   var port = cfg.Section("mysql").Key("port").String()
   var user = cfg.Section("mysql").Key("user").String()
   var passwd = cfg.Section("mysql").Key("passwd").String()
   var database = cfg.Section("mysql").Key("database").String()

   var build strings.Builder
   build.WriteString(user)
   build.WriteString(":")
   build.WriteString(passwd)
   build.WriteString("@tcp(")
   build.WriteString(ip)
   build.WriteString(":")
   build.WriteString(port)
   build.WriteString(")/")
   build.WriteString(database)
   build.WriteString("?charset=utf8")
   conn := build.String()

   //db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/vms?charset=utf8")
   db, err := sql.Open("mysql", conn)
   if err != nil {
       return nil
       }
   return db
}
