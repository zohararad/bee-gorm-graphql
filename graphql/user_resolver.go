package graphql

import (
  "github.com/zohararad/bee-gorm-graphql/models"
  "github.com/neelance/graphql-go"
)

type userResolver struct {
  entity *models.User
}

func (r *userResolver) ID() graphql.ID {
  return graphql.ID(r.entity.ID)
}

func (r *userResolver) FirstName() string {
  return r.entity.FirstName
}

func (r *userResolver) LastName() string {
  return r.entity.LastName
}

func (r *userResolver) Email() string {
  return r.entity.Email
}

func (r *userResolver) Password() string {
  return "********"
}

func ResolveUsers() (result []*userResolver) {
  users, err := models.AllUsers()
  if err != nil {
    return
  }
  for _, user := range users {
    result = append(result, &userResolver{entity: user})
  }
  return
}

func ResolveUser(id string) (result *userResolver) {
  user, err := models.FindUser(id)
  if err != nil {
    return nil
  }
  result = &userResolver{entity: user}
  return
}

func ResolveCreateUser(firstName string, lastName string, email string, password string) (result *userResolver) {
  user, err := models.CreateUser(firstName, lastName, email, password)
  if err != nil {
    return nil
  }
  result = &userResolver{entity: user}
  return
}

func ResolveUpdateUser(id string, firstName *string, lastName *string, email *string, password *string) (result *userResolver) {
  user, err := models.UpdateUser(id, firstName, lastName, email, password)
  if err != nil {
    return nil
  }
  result = &userResolver{entity: user}
  return
}

func ResolveDeleteUser(id string) *string {
  user, err := models.DeleteUser(id)
  if err != nil {
    return nil
  }
  return &user.ID
}