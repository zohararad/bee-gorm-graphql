package graphql

import (
  "context"
  "github.com/neelance/graphql-go"
)

type Resolver struct{}

func (r *Resolver) getCurrentUser(ctx context.Context) map[string]string {
  return ctx.Value("current_user").(map[string]string)
}

func (r *Resolver) Users(ctx context.Context) []*userResolver {
  return ResolveUsers()
}

func (r *Resolver) User(ctx context.Context, args *struct{ Id graphql.ID }) *userResolver {
  return ResolveUser(string(args.Id))
}

func (r *Resolver) CreateUser(ctx context.Context, args *struct{
  FirstName string; LastName string; Email string; Password string }) *userResolver {
  return ResolveCreateUser(args.FirstName, args.LastName, args.Email, args.Password)
}

func (r *Resolver) UpdateUser(ctx context.Context, args *struct{
  Id graphql.ID; FirstName *string; LastName *string; Email *string; Password *string }) *userResolver {
  return ResolveUpdateUser(string(args.Id), args.FirstName, args.LastName, args.Email, args.Password)
}

func (r *Resolver) DeleteUser(ctx context.Context, args *struct{ Id graphql.ID }) *graphql.ID {
  id := ResolveDeleteUser(string(args.Id))
  if id == nil {
    return nil
  }
  return &args.Id
}
