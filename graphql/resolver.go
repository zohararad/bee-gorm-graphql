package graphql

import "context"

type Resolver struct{}

func (r *Resolver) getCurrentUser(ctx context.Context) map[string]string {
  return ctx.Value("current_user").(map[string]string)
}
