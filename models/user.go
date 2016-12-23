package models

import (
  "github.com/zohararad/bee-gorm-graphql/db"
  "fmt"
  "crypto/sha256"
  "github.com/astaxie/beego"
)

type User struct {
  ID string                 `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
  FirstName string          `gorm:"type:varchar(128);not null;"`
  LastName string           `gorm:"type:varchar(128);not null;"`
  Email string              `gorm:"type:varchar(128);not null;unique_index"`
  EncryptedPassword string  `gorm:"type:varchar(128);not null;unique_index"`
}

func (u *User) AsMap() map[string]string {
  user := make(map[string]string)
  user["id"] = u.ID
  user["name"] = fmt.Sprintf("%s %s", u.FirstName, u.LastName)
  user["email"] = u.Email
  return user
}

func init() {
  db.RegisterMigration(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`)
  db.RegisterModel(&User{})
}

func UserLogin(email string, password string) (user *User, err error) {
  user = &User{}
  saltedPassword := fmt.Sprintf("%s_%s", password, beego.AppConfig.String("salt"))
  encryptedPassword := fmt.Sprintf("%x", sha256.Sum256([]byte(saltedPassword)))
  db := db.Conn.Where("email = ? AND encrypted_password = ?", email, encryptedPassword).First(user).Scan(user)
  err = db.Error
  if err != nil {
    return nil, err
  }
  return
}