package db

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
  "fmt"
)

var Conn *gorm.DB

func Connect(host string, database string, user string, pass string) (db *gorm.DB, err error) {
  dbConnString := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", host, user, database, pass)
  db, err = gorm.Open("postgres", dbConnString)
  Conn = db
  return
}