package main

import (
  _ "github.com/zohararad/bee-gorm-graphql/routers"
  _ "github.com/zohararad/bee-gorm-graphql/models" // import models package so models initialize and register
  "github.com/astaxie/beego"
  "github.com/zohararad/bee-gorm-graphql/db"
  "github.com/jinzhu/gorm"
)

func main() {
  conn := getDBConnection()
  defer conn.Close()
  beego.Run()
}

func getDBConnection() *gorm.DB {
  dbHost := beego.AppConfig.String("dbHost")
  dbUser := beego.AppConfig.String("dbUser")
  dbName := beego.AppConfig.String("dbName")
  dbPass := beego.AppConfig.String("dbPass")
  conn, err := db.Connect(dbHost, dbName, dbUser, dbPass)
  if err != nil {
    panic(err.Error())
  }
  db.AutoMigrate()
  return conn
}

