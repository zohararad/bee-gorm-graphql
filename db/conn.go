package db

import (
  "github.com/jinzhu/gorm"
  "fmt"
)

var Conn *gorm.DB

func Connect(host string, database string, user string, pass string) (*gorm.DB, error) {
  dbConnString := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", host, user, database, pass)
  Conn, err := gorm.Open("postgres", dbConnString)
  return Conn, err
}