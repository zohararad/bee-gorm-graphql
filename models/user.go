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
  EncryptedPassword string  `gorm:"type:varchar(128);not null;index"`
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
  encryptedPassword := userEncrypedPassword(password)
  db := db.Conn.Where("email = ? AND encrypted_password = ?", email, encryptedPassword).First(user).Scan(user)
  err = db.Error
  if err != nil {
    return nil, err
  }
  return
}

func userEncrypedPassword(password string) (encryptedPassword string) {
  saltedPassword := fmt.Sprintf("%s_%s", password, beego.AppConfig.String("salt"))
  encryptedPassword = fmt.Sprintf("%x", sha256.Sum256([]byte(saltedPassword)))
  return
}

func AllUsers() (users []User, err error) {
  res := db.Conn.Debug().Find(&users)
  err = res.Error
  if err != nil {
    beego.BeeLogger.Error("Error finding users: %v", res.Error.Error())
  }
  return
}

func FindUser(id string) (user User, err error) {
  res := db.Conn.Debug().First(&user, "id=?", id)
  err = res.Error
  if err != nil {
    beego.BeeLogger.Error("Error finding user with id %s: %v", id, res.Error.Error())
  }
  return
}

func CreateUser(firstName string, lastName string, email string, password string) (user User, err error) {
  user.FirstName = firstName
  user.LastName = lastName
  user.Email = email
  user.EncryptedPassword = userEncrypedPassword(password)
  res := db.Conn.Debug().Create(&user)
  err = res.Error
  return
}

func UpdateUser(id string, firstName *string, lastName *string, email *string, password *string) (User, error) {
  user := User{}
  res := db.Conn.Debug().First(&user, "id=?", id)
  if res.Error != nil {
    beego.BeeLogger.Error("Error finding user with id %s: %v", id, res.Error.Error())
  }
  if firstName != nil {
    user.FirstName = *firstName
  }
  if lastName != nil {
    user.LastName = *lastName
  }
  if email != nil {
    user.Email = *email
  }
  if password != nil {
    user.EncryptedPassword = userEncrypedPassword(*password)
  }
  res = db.Conn.Debug().Save(&user)
  err := res.Error
  return user, err
}

func DeleteUser(id string) (user User, err error) {
  user.ID = id
  res := db.Conn.Debug().Delete(&user)
  err = res.Error
  return
}